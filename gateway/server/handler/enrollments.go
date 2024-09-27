package handler

import (
	"context"
	"fmt"
	"gateway-service/config"
	"gateway-service/model"
	"gateway-service/pb"
	"gateway-service/server/middlewares"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"strconv"
	"time"
)

type EnrollmentHandler struct {
	enrollmentService pb.EnrollmentServiceClient
	userService       pb.UserServiceClient
	courseService     pb.CourseServiceClient
	userLog           pb.UserActivityLogServiceClient
	PaymentsHandler   PaymentsHandler
}

func NewEnrollmentHandler(enrollmentService pb.EnrollmentServiceClient, userService pb.UserServiceClient, courseService pb.CourseServiceClient) *EnrollmentHandler {
	return &EnrollmentHandler{
		enrollmentService: enrollmentService,
		userService:       userService,
		courseService:     courseService,
	}
}

func (h *EnrollmentHandler) CreateEnrollment(c echo.Context) error {
	// Memulai handler
	log.Println("[DEBUG] Starting CreateEnrollment handler")

	// Bind request body ke struct
	createEnrollmentRequest := new(model.EnrollmentInput)
	if err := c.Bind(createEnrollmentRequest); err != nil {
		log.Printf("[ERROR] Failed to bind request body: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Cek PaymentStatus dan beri default jika kosong
	if createEnrollmentRequest.PaymentStatus == "" {
		log.Println("[DEBUG] Setting default PaymentStatus to 'Pending'")
		createEnrollmentRequest.PaymentStatus = "Pending"
	}
	log.Printf("[DEBUG] Received request body: %+v\n", createEnrollmentRequest)

	// Ambil informasi user dari JWT token
	log.Println("[DEBUG] Binding JWT claims")
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		log.Println("[ERROR] User token is nil")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User not authorized"})
	}
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	if claims == nil || claims.UserID == "" {
		log.Println("[ERROR] Invalid JWT claims")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	// Log UserID dari klaim JWT
	studentId := claims.UserID
	email := claims.Email
	log.Printf("[DEBUG] JWT claims valid - StudentID: %s, Email: %s\n", studentId, email)
	createEnrollmentRequest.StudentID = studentId

	// Validasi input
	log.Println("[DEBUG] Validating input fields")
	if err := config.Validator.Struct(createEnrollmentRequest); err != nil {
		log.Printf("[ERROR] Validation errors: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Buat context dengan timeout
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	// Mendapatkan detail course
	log.Printf("[DEBUG] Fetching course details for CourseID: %d\n", createEnrollmentRequest.CourseID)
	if h.courseService == nil {
		log.Println("[ERROR] courseService is not initialized")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Course service is not available"})
	}

	courseResp, err := h.courseService.GetCourseById(ctx, &pb.GetCourseByIdRequest{
		Id: createEnrollmentRequest.CourseID,
	})
	if err != nil {
		log.Printf("[ERROR] Failed to retrieve course information: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve course information"})
	}
	courseName := courseResp.Title
	coursePrice := courseResp.Price
	log.Printf("[DEBUG] Retrieved course details - Name: %s, Price: %.2f\n", courseName, coursePrice)

	// Buat gRPC request untuk create enrollment
	enrollmentReq := &pb.CreateEnrollmentRequest{
		StudentId:     createEnrollmentRequest.StudentID,
		CourseId:      createEnrollmentRequest.CourseID,
		PaymentStatus: createEnrollmentRequest.PaymentStatus,
	}
	log.Printf("[DEBUG] gRPC enrollment request: %+v\n", enrollmentReq)

	// Panggil gRPC service untuk membuat enrollment
	log.Println("[DEBUG] Calling gRPC CreateEnrollment service")
	enrollmentResp, err := h.enrollmentService.CreateEnrollment(ctx, enrollmentReq)
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok && grpcErr.Code() == codes.AlreadyExists {
			log.Printf("[INFO] Enrollment already exists for student %s: %v\n", studentId, grpcErr)
			return c.JSON(http.StatusConflict, map[string]string{"error": grpcErr.Message()})
		}
		log.Printf("[ERROR] Failed to create enrollment via gRPC: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create enrollment"})
	}
	log.Printf("[DEBUG] Enrollment created successfully with ID: %d\n", enrollmentResp.Data.Id)

	// Generate invoice dan kirim email pembayaran
	log.Println("[DEBUG] Generating invoice and sending payment email")
	invoiceURL, externalId, price, err := h.PaymentsHandler.CreateInvoiceAndSendEmailPayment(studentId, email, courseName, coursePrice)
	if err != nil {
		log.Printf("[ERROR] Failed to create invoice and send payment email: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	log.Printf("[DEBUG] Invoice created - URL: %s, ExternalID: %s, AmountDue: %.2f\n", invoiceURL, externalId, price)

	// Log aktivitas user untuk enrollment
	log.Println("[DEBUG] Logging user activity")
	_, err = h.userLog.CreateUserActivityLog(ctx, &pb.CreateUserActivityLogRequest{
		UserId:            studentId,
		CourseId:          createEnrollmentRequest.CourseID,
		ActivityType:      "Enrollment in course " + courseName,
		ActivityTimestamp: timestamppb.New(time.Now()),
	})
	if err != nil {
		log.Printf("[ERROR] Failed to log user activity: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to log user activity"})
	}
	log.Println("[DEBUG] User activity logged successfully")

	// Membuat respon sukses
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"message": "Enrollment created successfully",
			"code":    http.StatusCreated,
			"status":  http.StatusText(http.StatusCreated),
		},
		"data": map[string]interface{}{
			"id":            enrollmentResp.Data.Id,
			"studentId":     enrollmentResp.Data.StudentId,
			"courseId":      enrollmentResp.Data.CourseId,
			"paymentStatus": enrollmentResp.Data.PaymentStatus,
			"enrolledAt":    enrollmentResp.Data.EnrolledAt,
			"createdAt":     enrollmentResp.Data.CreatedAt,
			"updatedAt":     enrollmentResp.Data.UpdatedAt,
			"invoice_url":   invoiceURL,
			"external_id":   externalId,
			"amount_due":    price,
		},
	}

	log.Println("[DEBUG] Returning success response")
	return c.JSON(http.StatusOK, response)
}

func (h *EnrollmentHandler) DeleteEnrollmentById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	studentId := claims.UserID

	idParam := c.Param("id")
	enrollmentID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid enrollment ID"})
	}
	req := &pb.DeleteEnrollmentByIdRequest{
		Id: uint32(enrollmentID),
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	resp, err := h.enrollmentService.DeleteEnrollmentById(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete enrollment"})
	}
	logReq := &pb.CreateUserActivityLogRequest{
		UserId:            studentId,
		ActivityType:      "delete enrollment ID : " + fmt.Sprint(enrollmentID),
		ActivityTimestamp: timestamppb.New(time.Now()),
	}
	_, err = h.userLog.CreateUserActivityLog(ctx, logReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to log user activity"})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *EnrollmentHandler) GetEnrollmentsByStudentId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	studentId := claims.UserID

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	enrollmentResp, err := h.enrollmentService.GetEnrollmentsByStudentId(ctx, &pb.GetEnrollmentsByStudentIdRequest{
		StudentId: studentId,
	})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok && grpcErr.Code() == codes.NotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": grpcErr.Message()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": grpcErr.Message()})
	}

	return c.JSON(http.StatusOK, enrollmentResp)
}

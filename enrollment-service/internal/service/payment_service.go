package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/config"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/proto/meta"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/proto/payments"
	pb "github.com/ghssni/Smartcy-LMS/enrollment-service/proto/payments"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/xendit/xendit-go/v6/invoice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PaymentService interface {
	pb.PaymentsServiceServer
	HandleWebhookHTTP(c echo.Context) error
}

type paymentService struct {
	pb.UnimplementedPaymentsServiceServer
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{
		repo: repo,
	}
}

func (s *paymentService) GetPaymentByEnrollmentId(ctx context.Context, req *payments.GetPaymentByEnrollmentIdRequest) (*payments.GetPaymentByEnrollmentIdResponse, error) {
	payment, err := s.repo.GetPaymentByEnrollmentId(ctx, strconv.Itoa(int(req.EnrollmentId)))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Payment not found for enrollment ID: %d", req.EnrollmentId)
	}

	return &payments.GetPaymentByEnrollmentIdResponse{
		Payments: &payments.Payments{
			ExternalId:             payment.ExternalID,
			UserId:                 payment.UserID,
			PaymentMethod:          payment.PaymentMethod,
			Status:                 payment.Status,
			MerchantName:           payment.MerchantName,
			Amount:                 float32(payment.Amount),
			PaidAmount:             float32(payment.PaidAmount),
			BankCode:               payment.BankCode,
			PaidAt:                 payment.PaidAt,
			PayerEmail:             payment.PayerEmail,
			Description:            payment.Description,
			AdjustedReceivedAmount: float32(payment.AdjustedReceivedAmount),
			FeesPaidAmount:         float32(payment.FeesPaidAmount),
			Updated:                timestamppb.New(payment.Updated),
			Created:                timestamppb.New(payment.Created),
			Currency:               payment.Currency,
			PaymentChannel:         payment.PaymentChannel,
			PaymentDestination:     payment.PaymentDestination,
		},
	}, nil
}

func (s *paymentService) HandleWebhook(ctx context.Context, req *payments.HandleWebhookRequest) (*payments.HandleWebhookResponse, error) {
	updatedPayment := models.Payments{
		ExternalID:             req.ExternalId,
		IsHigh:                 false,
		PaymentMethod:          req.PaymentMethod,
		Status:                 req.Status,
		MerchantName:           req.MerchantName,
		Amount:                 float64(req.Amount),
		PaidAmount:             float64(req.Amount),
		BankCode:               req.BankCode,
		PaidAt:                 req.PaidAt,
		PayerEmail:             req.Email,
		Description:            req.Description,
		AdjustedReceivedAmount: float64(req.Amount),
		FeesPaidAmount:         float64(req.Amount),
		Updated:                time.Now(),
		Created:                time.Now(),
		Currency:               req.Currency,
		PaymentChannel:         req.PaymentChannel,
		PaymentDestination:     req.PaymentDestination,
		InvoiceUrl:             req.InvoiceUrl,
	}

	newInvoice, err := s.repo.UpdatePaymentStatusByWebhook(ctx, req.ExternalId, updatedPayment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update payment status: %v", err)
	}

	//Send email to student if payment is successful
	//if req.Status == "PAID" {
	//	err := config.SendEmailSuccess(newInvoice.PayerEmail, newInvoice.Description)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	return &payments.HandleWebhookResponse{
		Meta: &meta.Meta{
			Message: "Payment status updated successfully",
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
		},
		Payments: &payments.Payments{
			ExternalId:             newInvoice.ExternalID,
			UserId:                 newInvoice.UserID,
			PaymentMethod:          newInvoice.PaymentMethod,
			Status:                 newInvoice.Status,
			MerchantName:           newInvoice.MerchantName,
			Amount:                 float32(newInvoice.Amount),
			PaidAmount:             float32(newInvoice.PaidAmount),
			BankCode:               newInvoice.BankCode,
			PaidAt:                 newInvoice.PaidAt,
			PayerEmail:             newInvoice.PayerEmail,
			Description:            newInvoice.Description,
			AdjustedReceivedAmount: float32(newInvoice.AdjustedReceivedAmount),
			FeesPaidAmount:         float32(newInvoice.FeesPaidAmount),
			Updated:                timestamppb.New(newInvoice.Updated),
			Created:                timestamppb.New(newInvoice.Created),
			Currency:               newInvoice.Currency,
			PaymentChannel:         newInvoice.PaymentChannel,
			PaymentDestination:     newInvoice.PaymentDestination,
			InvoiceUrl:             newInvoice.InvoiceUrl,
		},
	}, nil
}

func (s *paymentService) HandleWebhookHTTP(c echo.Context) error {
	if err := verifyXenditWebhook(c); err != nil {
		logrus.Errorf("Webhook verification failed: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized webhook: " + err.Error(),
		})
	}

	bodyBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		logrus.Errorf("Failed to read request body: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to read request body",
		})
	}

	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	webhookRequest := new(pb.HandleWebhookRequest)
	if err := c.Bind(webhookRequest); err != nil {
		logrus.Errorf("Failed to bind webhook request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Ignore webhook for expired transaction
	if webhookRequest.Status == "EXPIRED" {
		logrus.Infof("Ignoring webhook for expired transaction: %s", webhookRequest.ExternalId)
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Webhook for expired transaction ignored",
		})
	}

	ctx := context.Background()
	webhookResponse, err := s.HandleWebhook(ctx, webhookRequest)
	if err != nil {
		logrus.Errorf("Error processing webhook: %v, WebhookRequest: %+v", err, webhookRequest)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error processing webhook: %v", err),
		})
	}

	return c.JSON(http.StatusOK, webhookResponse)
}

func CreateInvoice(externalID, email, courseName string, price float64) (string, error) {
	xenditClient := config.XenditClient
	description := "Payment for " + courseName
	ctx := context.Background()
	// Create invoice
	resp, httpResponse, err := xenditClient.InvoiceApi.CreateInvoice(ctx).
		CreateInvoiceRequest(invoice.CreateInvoiceRequest{
			ExternalId:  externalID,
			PayerEmail:  &email,
			Description: &description,
			Amount:      price,
		}).
		Execute()

	if err != nil {
		return "", fmt.Errorf("error creating invoice: %v", err)
	}

	if httpResponse.StatusCode != 200 {
		return "", fmt.Errorf("error creating invoice: %v", httpResponse.Body)
	}

	return resp.InvoiceUrl, nil

}

// CreateInvoiceAndSendEmailPayment creates an invoice and sends an email to the student with the payment link
func CreateInvoiceAndSendEmailPayment(studentId, email, courseName string, price float64) (string, string, float64, error) {
	externalId := fmt.Sprintf("invoice_%s_%d", studentId, time.Now().Unix())
	invoiceURL, err := CreateInvoice(externalId, email, courseName, price)
	if err != nil {
		return "", "", 0, fmt.Errorf("error creating invoice: %v", err)
	}

	//err = config.SendEmailPayment(email, courseName, invoiceURL)
	//if err != nil {
	//	return "", "", 0, fmt.Errorf("error creating invoice: %v", err)
	//}

	return invoiceURL, externalId, price, nil
}

// verifyXenditWebhook verifies the Xendit webhook signature
func verifyXenditWebhook(c echo.Context) error {
	callbackToken := c.Request().Header.Get("X-CALLBACK-TOKEN")
	expectedToken := os.Getenv("XENDIT_CALLBACK_TOKEN")

	if callbackToken != expectedToken {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid callback token")
	}
	return nil
}

// UpdateExpiredPaymentStatus updates the status of expired payments
func (s *paymentService) UpdateExpiredPaymentStatus(ctx context.Context, req *pb.UpdateExpiredPaymentStatusRequest) (*pb.UpdateExpiredPaymentStatusResponse, error) {
	if err := s.repo.UpdateExpiredPaymentStatus(); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update expired payment: %v", err)
	}

	return &pb.UpdateExpiredPaymentStatusResponse{
		Meta: &meta.Meta{
			Message: "Expired payment status updated successfully",
			Code:    int32(codes.OK),
			Status:  codes.OK.String(),
		},
	}, nil
}

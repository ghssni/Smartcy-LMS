package handler

import (
	"bytes"
	"context"
	"fmt"
	"gateway-service/config"
	"gateway-service/pb"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/xendit/xendit-go/v6/invoice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type PaymentsHandler struct {
	paymentService pb.PaymentsServiceClient
	userService    pb.UserServiceClient
	emailService   pb.EmailServiceClient
}

func NewPaymentsHandler(enrollmentService pb.PaymentsServiceClient, userService pb.UserServiceClient, client pb.EmailServiceClient) *PaymentsHandler {
	return &PaymentsHandler{
		paymentService: enrollmentService,
		userService:    userService,
		emailService:   client,
	}
}

func (h *PaymentsHandler) GetPaymentByEnrollmentId(c echo.Context) error {
	enrollmentIdStr := c.Param("enrollmentId")
	enrollmentId, err := strconv.ParseUint(enrollmentIdStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid enrollment ID"})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	paymentResp, err := h.paymentService.GetPaymentByEnrollmentId(ctx, &pb.GetPaymentByEnrollmentIdRequest{
		EnrollmentId: uint32(enrollmentId),
	})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok && grpcErr.Code() == codes.NotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": grpcErr.Message()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve payment"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"external_id":              paymentResp.Payments.ExternalId,
		"user_id":                  paymentResp.Payments.UserId,
		"payment_method":           paymentResp.Payments.PaymentMethod,
		"status":                   paymentResp.Payments.Status,
		"merchant_name":            paymentResp.Payments.MerchantName,
		"amount":                   paymentResp.Payments.Amount,
		"paid_amount":              paymentResp.Payments.PaidAmount,
		"bank_code":                paymentResp.Payments.BankCode,
		"paid_at":                  paymentResp.Payments.PaidAt,
		"payer_email":              paymentResp.Payments.PayerEmail,
		"description":              paymentResp.Payments.Description,
		"adjusted_received_amount": paymentResp.Payments.AdjustedReceivedAmount,
		"fees_paid_amount":         paymentResp.Payments.FeesPaidAmount,
		"updated":                  paymentResp.Payments.Updated,
		"created":                  paymentResp.Payments.Created,
		"currency":                 paymentResp.Payments.Currency,
		"payment_channel":          paymentResp.Payments.PaymentChannel,
		"payment_destination":      paymentResp.Payments.PaymentDestination,
	})
}

// CreateInvoice creates an invoice with Xendit API
func CreateInvoice(externalID, email, courseName string, price float64) (string, error) {
	xenditClient := config.XenditClient
	description := "Payment for " + courseName
	ctx := context.Background()

	// Create invoice
	resp, _, err := xenditClient.InvoiceApi.CreateInvoice(ctx).
		CreateInvoiceRequest(invoice.CreateInvoiceRequest{
			ExternalId:  externalID,
			PayerEmail:  &email,
			Description: &description,
			Amount:      price,
		}).Execute()

	if err != nil {
		return "", fmt.Errorf("error creating invoice: %v", err)
	}

	return resp.InvoiceUrl, nil
}

// CreateInvoiceAndSendEmailPayment creates an invoice and sends an email to the student with the payment link
func (h *PaymentsHandler) CreateInvoiceAndSendEmailPayment(studentId, email, courseName string, price float64) (string, string, float64, error) {
	externalId := fmt.Sprintf("invoice_%s_%d", studentId, time.Now().Unix())

	// Create Invoice
	invoiceURL, err := CreateInvoice(externalId, email, courseName, price)
	if err != nil {
		return "", "", 0, fmt.Errorf("error creating invoice: %v", err)
	}

	// Prepare email request
	emailReq := &pb.SendPaymentDueEmailRequest{
		Email:       email,
		UserId:      studentId,
		CourseName:  courseName,
		PaymentLink: invoiceURL,
	}

	// Send email
	emailResp, err := h.emailService.SendPaymentDueEmail(context.Background(), emailReq)
	if err != nil {
		return "", "", 0, fmt.Errorf("error sending payment email: %v", err)
	}

	// Check if email sending was successful
	if !emailResp.Success {
		return "", "", 0, fmt.Errorf("failed to send payment email: %v", emailResp.Meta.Message)
	}

	return invoiceURL, externalId, price, nil
}

// verifyXenditWebhook verifies the Xendit webhook signature
func verifyXenditWebhook(c echo.Context) error {
	callbackToken := c.Request().Header.Get("X-CALLBACK-TOKEN")
	expectedToken := "4ya7xNWlhClmYznAwYBUeQxQviieT9gzhKacK00zZZxGM4yP"

	if callbackToken != expectedToken {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid callback token")
	}
	return nil
}

func (h *PaymentsHandler) HandleWebhook(c echo.Context) error {
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

	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()
	if _, err := h.paymentService.HandleWebhook(ctx, webhookRequest); err != nil {
		logrus.Errorf("Failed to handle webhook: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to handle webhook",
		})
	}

	// if status paid
	if webhookRequest.Status == "PAID" {
		emailReq := &pb.SendPaymentSuccessEmailRequest{
			Email:       webhookRequest.Email,
			CourseName:  webhookRequest.Description,
			Amount:      webhookRequest.Amount,
			Description: webhookRequest.Description,
			InvoiceId:   webhookRequest.ExternalId,
		}
		_, err := h.emailService.SendPaymentSuccessEmail(ctx, emailReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send payment success email"})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Webhook received successfully",
	})
}

package service

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/models"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"net/http"
	"strconv"
	"time"
)

type PaymentService interface {
	pb.PaymentsServiceServer
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

func (s *paymentService) GetPaymentByEnrollmentId(ctx context.Context, req *pb.GetPaymentByEnrollmentIdRequest) (*pb.GetPaymentByEnrollmentIdResponse, error) {
	payment, err := s.repo.GetPaymentByEnrollmentId(ctx, strconv.Itoa(int(req.EnrollmentId)))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Payment not found for enrollment ID: %d", req.EnrollmentId)
	}

	return &pb.GetPaymentByEnrollmentIdResponse{
		Payments: &pb.Payments{
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

func (s *paymentService) HandleWebhook(ctx context.Context, req *pb.HandleWebhookRequest) (*pb.HandleWebhookResponse, error) {

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

	return &pb.HandleWebhookResponse{
		Meta: &pb.MetaPayments{
			Message: "Payment status updated successfully",
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
		},
		Payments: &pb.Payments{
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

// UpdateExpiredPaymentStatus updates the status of expired pb
func (s *paymentService) UpdateExpiredPaymentStatus(ctx context.Context, req *pb.UpdateExpiredPaymentStatusRequest) (*pb.UpdateExpiredPaymentStatusResponse, error) {
	if err := s.repo.UpdateExpiredPaymentStatus(); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update expired payment: %v", err)
	}

	return &pb.UpdateExpiredPaymentStatusResponse{
		Meta: &pb.MetaPayments{
			Message: "Expired payment status updated successfully",
			Code:    uint32(codes.OK),
			Status:  codes.OK.String(),
		},
	}, nil
}

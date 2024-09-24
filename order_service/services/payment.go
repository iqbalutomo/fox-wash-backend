package services

import (
	"context"
	"order_service/models"

	"github.com/xendit/xendit-go/v3"
	"github.com/xendit/xendit-go/v3/invoice"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PaymentService interface {
	CreateInvoice(externalID primitive.ObjectID, subtotal float32) (models.Payment, error)
}

type XenditClient struct {
	client *xendit.APIClient
}

func NewPaymentService(key string) PaymentService {
	return &XenditClient{
		client: xendit.NewClient(key),
	}
}

func (x *XenditClient) CreateInvoice(externalID primitive.ObjectID, subtotal float32) (models.Payment, error) {
	createInvoiceReq := *invoice.NewCreateInvoiceRequest(externalID.Hex(), subtotal)
	resp, _, err := x.client.InvoiceApi.CreateInvoice(context.Background()).CreateInvoiceRequest(createInvoiceReq).Execute()
	if err != nil {
		return models.Payment{}, status.Error(codes.Internal, err.Error())
	}

	paymentData := models.Payment{
		InvoiceID:  *resp.Id,
		InvoiceURL: resp.InvoiceUrl,
		Total:      resp.Amount,
		Method:     "pending",
		Status:     resp.Status.String(),
	}

	return paymentData, nil
}

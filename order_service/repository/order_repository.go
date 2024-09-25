package repository

import (
	"context"
	"order_service/models"
	"order_service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Order interface {
	CreateOrder(ctx context.Context, data *models.Order) error
	UpdateOrderPaymentStatus(ctx context.Context, invoiceID, status, method, completeAt string) error
}

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) Order {
	return &OrderRepository{collection}
}

func (o *OrderRepository) CreateOrder(ctx context.Context, data *models.Order) error {
	res, err := o.collection.InsertOne(ctx, data)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	data.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (o *OrderRepository) UpdateOrderPaymentStatus(ctx context.Context, invoiceID, status, method, completeAt string) error {
	filter := bson.D{{Key: "payment.invoice_id", Value: invoiceID}}
	updateData := bson.M{"$set": bson.M{
		"payment.status":       status,
		"payment.method":       method,
		"payment.completed_at": completeAt,
		"status":               utils.OrderStatusWasherPreparing,
	}}
	if err := o.UpdateWithFilter(ctx, filter, updateData); err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) UpdateWithFilter(ctx context.Context, field bson.D, data bson.M) error {
	res := o.collection.FindOneAndUpdate(ctx, field, data)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return status.Error(codes.NotFound, err.Error())
		}

		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

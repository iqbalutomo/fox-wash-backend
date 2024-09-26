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
	FindByID(ctx context.Context, orderID string) (models.Order, error)
	FindWasherAllOrders(ctx context.Context, washerID uint) ([]models.Order, error)
	FindWasherCurrentOrder(ctx context.Context, washerID uint) (models.Order, error)
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

func (o *OrderRepository) FindByID(ctx context.Context, orderID string) (models.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return models.Order{}, status.Error(codes.InvalidArgument, err.Error())
	}

	res := o.collection.FindOne(ctx, bson.M{"_id": objectID})
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Order{}, status.Error(codes.NotFound, err.Error())
		}

		return models.Order{}, status.Error(codes.Internal, err.Error())
	}

	var order models.Order
	if err := res.Decode(&order); err != nil {
		return models.Order{}, status.Error(codes.Internal, err.Error())
	}

	return order, nil
}

func (o *OrderRepository) FindWasherAllOrders(ctx context.Context, washerID uint) ([]models.Order, error) {
	filter := bson.D{{
		Key:   "washer.id",
		Value: washerID,
	}}

	orders, err := o.FindWithFilter(ctx, filter)
	if err != nil {
		return []models.Order{}, status.Error(codes.Internal, err.Error())
	}

	return orders, nil
}

func (o *OrderRepository) FindWasherCurrentOrder(ctx context.Context, washerID uint) (models.Order, error) {
	filter := bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "washer.id", Value: washerID}},
			bson.D{{Key: "status", Value: utils.OrderStatusWasherPreparing}},
			bson.D{{Key: "payment.status", Value: "PAID"}},
		}},
	}

	res := o.collection.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Order{}, status.Error(codes.NotFound, err.Error())
		}

		return models.Order{}, status.Error(codes.Internal, err.Error())
	}

	var order models.Order
	if err := res.Decode(&order); err != nil {
		return models.Order{}, status.Error(codes.Internal, err.Error())
	}

	return order, nil
}

func (o *OrderRepository) FindWithFilter(ctx context.Context, filter bson.D) ([]models.Order, error) {
	res, err := o.collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []models.Order{}, status.Error(codes.NotFound, err.Error())
		}

		return []models.Order{}, status.Error(codes.Internal, err.Error())
	}

	orders := []models.Order{}
	if err := res.All(ctx, &orders); err != nil {
		return []models.Order{}, status.Error(codes.Internal, err.Error())
	}

	return orders, nil
}

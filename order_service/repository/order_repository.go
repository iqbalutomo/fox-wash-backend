package repository

import (
	"context"
	"order_service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Order interface {
	CreateOrder(ctx context.Context, data *models.Order) error
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

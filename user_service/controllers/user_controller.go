package controllers

import (
	"context"
	"encoding/json"
	"time"
	"user_service/dto"
	"user_service/helpers"
	"user_service/models"
	"user_service/pb"
	"user_service/repository"
	"user_service/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUserServer
	repo repository.User
	mb   services.MessageBroker
}

func NewUserController(repo repository.User, mb services.MessageBroker) *Server {
	return &Server{
		repo: repo,
		mb:   mb,
	}
}

func (s *Server) Register(ctx context.Context, data *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := helpers.HashingPassword(data.Password)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := s.repo.CreateUser(&newUser); err != nil {
		return nil, err
	}

	dataJsonRequest := dto.UserMessageBroker{
		ID:    newUser.ID,
		Name:  newUser.FirstName,
		Email: newUser.Email,
	}

	dataJson, err := json.Marshal(dataJsonRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.mb.PublishMessage(dataJson); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &pb.RegisterResponse{
		UserId:    uint32(newUser.ID),
		CreatedAt: newUser.CreatedAt,
	}

	return response, nil
}

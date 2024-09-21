package controllers

import (
	"context"
	"time"
	"user_service/helpers"
	"user_service/models"
	"user_service/pb"
	"user_service/repository"
)

type Server struct {
	pb.UnimplementedUserServer
	repo repository.User
}

func NewUserController(repo repository.User) *Server {
	return &Server{repo: repo}
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

	response := &pb.RegisterResponse{
		UserId:    uint32(newUser.ID),
		CreatedAt: newUser.CreatedAt,
	}

	return response, nil
}

package controllers

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"user_service/dto"
	"user_service/helpers"
	"user_service/models"
	"user_service/pb"
	"user_service/repository"
	"user_service/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
	newUser := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		RoleID:    data.RoleId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := s.repo.CreateUser(&newUser); err != nil {
		return nil, err
	}

	tokenVerify, err := helpers.GenerateTokenVerify()
	if err != nil {
		log.Fatal(err)
	}

	verificationData := models.EmailVerification{
		UserID: newUser.ID,
		Token:  tokenVerify,
	}

	if err := s.repo.AddToken(&verificationData); err != nil {
		return nil, err
	}

	dataJsonRequest := dto.UserMessageBroker{
		ID:    newUser.ID,
		Name:  newUser.FirstName,
		Email: newUser.Email,
		Token: verificationData.Token,
	}

	dataJson, err := json.Marshal(dataJsonRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.mb.PublishMessageVerification(dataJson); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &pb.RegisterResponse{
		UserId:    uint32(newUser.ID),
		CreatedAt: newUser.CreatedAt,
	}

	return response, nil
}

func (s *Server) VerifyNewUser(ctx context.Context, data *pb.UserCredential) (*emptypb.Empty, error) {
	if err := s.repo.VerifyNewUser(data.Id, data.Token); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetUser(ctx context.Context, data *pb.EmailRequest) (*pb.UserData, error) {
	user, err := s.repo.GetUser(data.Email)
	if err != nil {
		return nil, err
	}

	userData := &pb.UserData{
		Id:         uint32(user.ID),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Password:   user.Password,
		Role:       user.Role,
		IsVerified: user.IsVerified,
	}

	return userData, nil
}

func (s *Server) CreateWasher(ctx context.Context, data *pb.WasherID) (*emptypb.Empty, error) {
	if err := s.repo.CreateWasher(data.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) WasherActivation(ctx context.Context, data *pb.EmailRequest) (*emptypb.Empty, error) {
	if err := s.repo.WasherActivation(data.Email); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetWasher(ctx context.Context, data *pb.WasherID) (*pb.WasherData, error) {
	washer, err := s.repo.GetWasher(int32(data.Id))
	if err != nil {
		return nil, err
	}

	washerData := &pb.WasherData{
		UserId:         uint32(washer.UserID),
		IsOnline:       washer.IsOnline,
		WasherStatusId: uint32(washer.WasherStatusID),
		IsActive:       washer.IsActive,
	}

	return washerData, nil
}

func (s *Server) GetAvailableWasher(ctx context.Context, data *emptypb.Empty) (*pb.WasherOrderData, error) {
	washer, err := s.repo.GetAvailableWasher()
	if err != nil {
		return nil, err
	}

	washerData := &pb.WasherOrderData{
		Id:     washer.ID,
		Name:   washer.Name,
		Status: washer.Status,
	}

	return washerData, nil
}

func (s *Server) SetWasherStatusOnline(ctx context.Context, data *pb.WasherID) (*emptypb.Empty, error) {
	washerID := data.Id

	if err := s.repo.SetWasherStatusOnline(washerID); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

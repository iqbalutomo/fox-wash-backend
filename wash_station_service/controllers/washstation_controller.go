package controllers

import (
	"context"
	"errors"
	"wash_station_service/pb"
	"wash_station_service/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedWashStationServer
	repo repository.WashStation
}

func NewWashStationController(repo repository.WashStation) *Server {
	return &Server{repo: repo}
}

func (s *Server) CreateWashPackage(ctx context.Context, data *pb.NewWashPackageData) (*pb.CreateWashPackageResponse, error) {
	return nil, errors.New("") // TODO: logic here
}

func (s *Server) FindAllWashPackages(ctx context.Context, empty *emptypb.Empty) (*pb.WashPackageCompactRepeated, error) {
	return nil, errors.New("") // TODO: logic here
}

func (s *Server) FindWashPackageByID(ctx context.Context, washPackageID *pb.WashPackageID) (*pb.WashPackageData, error) {
	return nil, errors.New("") // TODO: logic here
}

func (s *Server) UpdateWashPackage(ctx context.Context, data *pb.UpdateWashPackageData) (*emptypb.Empty, error) {
	return nil, errors.New("") // TODO: logic here
}

func (s *Server) DeleteWashPackage(ctx context.Context, washPackageID *pb.WashPackageID) (*emptypb.Empty, error) {
	return nil, errors.New("") // TODO: logic here
}

package controllers

import (
	"context"
	"errors"
	"time"
	"wash_station_service/models"
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
	washPackageData := models.Wash{
		Name:      data.Name,
		Category:  data.Category,
		Price:     float64(data.Price),
		CreatedBy: data.CreatedBy,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := s.repo.CreateWashPackage(&washPackageData); err != nil {
		return nil, err
	}

	return &pb.CreateWashPackageResponse{Id: washPackageData.ID}, nil
}

func (s *Server) FindAllWashPackages(ctx context.Context, empty *emptypb.Empty) (*pb.WashPackageCompactRepeated, error) {
	washPackages, err := s.repo.FindAllWashPackages()
	if err != nil {
		return nil, err
	}

	var pbWashPackages []*pb.WashPackageCompact
	for _, wash := range washPackages {
		pbWashPackage := &pb.WashPackageCompact{
			Id:       wash.ID,
			Name:     wash.Name,
			Category: wash.Category,
			Price:    float32(wash.Price),
		}

		pbWashPackages = append(pbWashPackages, pbWashPackage)
	}

	return &pb.WashPackageCompactRepeated{
		WashPackages: pbWashPackages,
	}, nil
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

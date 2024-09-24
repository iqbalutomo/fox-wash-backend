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

//detailing package controller

func (s *Server) CreateDetailingPackage(ctx context.Context, data *pb.NewDetailingPackageData) (*pb.CreateDetailingPackageResponse, error) {
	detailingPackageData := models.Detailing{
		Name:        data.Name,
		Description: data.Description,
		Price:       float64(data.Price),
		CreatedBy:   data.CreatedBy,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := s.repo.CreateDetailingPackage(&detailingPackageData); err != nil {
		return nil, err
	}
	return &pb.CreateDetailingPackageResponse{
		Id: detailingPackageData.ID,
	}, nil
}

func (s *Server) FindAllDetailingPackages(ctx context.Context, empty *emptypb.Empty) (*pb.DetailingPackageCompactRepeated, error) {
	var detailingPackages []*pb.DetailingPackageCompact
	detailingPackageData, err := s.repo.FindAllDetailingPackages()
	if err != nil {
		return nil, err
	}
	for _, data := range detailingPackageData {
		detailingPackages = append(detailingPackages, &pb.DetailingPackageCompact{
			Id:          data.ID,
			Name:        data.Name,
			Description: data.Description,
			Price:       float32(data.Price),
		})
	}
	return &pb.DetailingPackageCompactRepeated{
		DetailingPackages: detailingPackages,
	}, nil
}

func (s *Server) FindDetailingPackageByID(ctx context.Context, detailingPackageID *pb.DetailingPackageID) (*pb.DetailingPackageData, error) {
	detailingPackageData, err := s.repo.FindDetailingPackageByID(detailingPackageID.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DetailingPackageData{
		Id:          detailingPackageData.ID,
		Name:        detailingPackageData.Name,
		Description: detailingPackageData.Description,
		Price:       float32(detailingPackageData.Price),
	}, nil
}

func (s *Server) UpdateDetailingPackage(ctx context.Context, data *pb.UpdateDetailingPackageData) (*emptypb.Empty, error) {
	if _, err := s.repo.FindDetailingPackageByID(data.Id); err != nil {
		return nil, err
	}
	if err := s.repo.UpdateDetailingPackage(data.Id, data); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteDetailingPackage(ctx context.Context, detailingPackageID *pb.DetailingPackageID) (*emptypb.Empty, error) {
	if _, err := s.repo.FindDetailingPackageByID(detailingPackageID.Id); err != nil {
		return nil, err
	}
	if err := s.repo.DeleteDetailingPackage(detailingPackageID.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

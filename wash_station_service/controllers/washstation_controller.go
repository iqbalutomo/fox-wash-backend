package controllers

import (
	"context"
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

func (s *Server) FindWashPackageByID(ctx context.Context, data *pb.WashPackageID) (*pb.WashPackageData, error) {
	result, err := s.repo.FindWashPackageByID(data.Id)
	if err != nil {
		return nil, err
	}

	washPackage := &pb.WashPackageData{
		Id:       result.ID,
		Name:     result.Name,
		Category: result.Category,
		Price:    float32(result.Price),
	}

	return washPackage, nil
}

func (s *Server) FindMultipleWashPackages(ctx context.Context, data *pb.WashPackageIDs) (*pb.WashPackageCompactRepeated, error) {
	washPackageIDs := data.GetIds()
	washPackages, err := s.repo.FindMultipleWashPackages(washPackageIDs)
	if err != nil {
		return nil, err
	}

	var washPackageDatas []*pb.WashPackageCompact
	for _, wash := range washPackages {
		washPackageDatas = append(washPackageDatas, &pb.WashPackageCompact{
			Id:       wash.Id,
			Name:     wash.Name,
			Category: wash.Category,
			Price:    wash.Price,
		})
	}

	response := &pb.WashPackageCompactRepeated{
		WashPackages: washPackageDatas,
	}

	return response, nil
}

func (s *Server) UpdateWashPackage(ctx context.Context, data *pb.UpdateWashPackageData) (*emptypb.Empty, error) {
	washPackageTmp, err := s.repo.FindWashPackageByID(data.Id)
	if err != nil {
		return nil, err
	}

	if err := s.repo.UpdateWashPackage(washPackageTmp.ID, data); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteWashPackage(ctx context.Context, data *pb.WashPackageID) (*emptypb.Empty, error) {
	if err := s.repo.DeleteWashPackage(data.Id); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
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

func (s *Server) FindMultipleDetailingPackages(ctx context.Context, data *pb.DetailingPackageIDs) (*pb.DetailingPackageCompactRepeated, error) {
	detailingPackageIDs := data.GetIds()
	detailingPackages, err := s.repo.FindMultipleDetailingPackages(detailingPackageIDs)
	if err != nil {
		return nil, err
	}

	var detailingPackageDatas []*pb.DetailingPackageCompact
	for _, wash := range detailingPackages {
		detailingPackageDatas = append(detailingPackageDatas, &pb.DetailingPackageCompact{
			Id:          wash.Id,
			Name:        wash.Name,
			Description: wash.Description,
			Price:       wash.Price,
		})
	}

	response := &pb.DetailingPackageCompactRepeated{
		DetailingPackages: detailingPackageDatas,
	}

	return response, nil
}

func (s *Server) UpdateDetailingPackage(ctx context.Context, data *pb.UpdateDetailingPackageData) (*emptypb.Empty, error) {
	if err := s.repo.UpdateDetailingPackage(data.Id, data); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteDetailingPackage(ctx context.Context, detailingPackageID *pb.DetailingPackageID) (*emptypb.Empty, error) {
	if err := s.repo.DeleteDetailingPackage(detailingPackageID.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

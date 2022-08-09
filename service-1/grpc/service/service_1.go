package service

import (
	"context"
	"service-1/config"
	pb "service-1/genproto/service_1"
	"service-1/pkg/logger"
	"service-1/storage"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type service_1 struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedService_1Server
}

func NewService1Service(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *service_1 {
	return &service_1{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *service_1) GetAll(ctx context.Context, req *emptypb.Empty) (*pb.Status, error) {
	resp, err := s.strg.Service_1_I().GetAll(ctx)
	if err != nil {
		s.log.Error("GetByIdAttribute", logger.Error(err))
		return nil, err
	}

	return &pb.Status{
		Status: resp.Status,
	}, nil
}

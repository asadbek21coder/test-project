package service

import (
	"context"

	"github.com/asadbek21coder/test-project/service-2/config"
	pb "github.com/asadbek21coder/test-project/service-2/genproto/service_2"
	"github.com/asadbek21coder/test-project/service-2/pkg/logger"
	"github.com/asadbek21coder/test-project/service-2/storage"
)

type service_2 struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedService_2Server
}

func NewService2Service(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *service_2 {
	return &service_2{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *service_2) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	resp, err := s.strg.Service_2_I().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllPosts", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *service_2) GetById(ctx context.Context, req *pb.Id) (*pb.Post, error) {
	resp, err := s.strg.Service_2_I().GetById(ctx, req)
	if err != nil {
		s.log.Error("GetByIdPost", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *service_2) Update(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	resp, err := s.strg.Service_2_I().Update(ctx, req)
	if err != nil {
		s.log.Error("UpdatePost", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *service_2) Delete(ctx context.Context, req *pb.Id) (*pb.Id, error) {
	resp, err := s.strg.Service_2_I().Delete(ctx, req)
	if err != nil {
		s.log.Error("DeletePost", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

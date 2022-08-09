package grpc

import (
	"service-1/config"
	"service-1/genproto/service_1"
	"service-1/grpc/service"
	"service-1/pkg/logger"
	"service-1/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()
	service_1.RegisterService_1Server(grpcServer, service.NewService1Service(cfg, log, strg))
	reflection.Register(grpcServer)
	return
}

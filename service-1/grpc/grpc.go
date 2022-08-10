package grpc

import (
	"github.com/asadbek21coder/test-project/service-1/config"
	"github.com/asadbek21coder/test-project/service-1/genproto/service_1"
	"github.com/asadbek21coder/test-project/service-1/grpc/service"
	"github.com/asadbek21coder/test-project/service-1/pkg/logger"
	"github.com/asadbek21coder/test-project/service-1/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()
	service_1.RegisterService_1Server(grpcServer, service.NewService1Service(cfg, log, strg))
	reflection.Register(grpcServer)
	return
}

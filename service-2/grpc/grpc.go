package grpc

import (
	"github.com/asadbek21coder/test-project/service-2/config"
	"github.com/asadbek21coder/test-project/service-2/genproto/service_2"
	"github.com/asadbek21coder/test-project/service-2/grpc/service"
	"github.com/asadbek21coder/test-project/service-2/pkg/logger"
	"github.com/asadbek21coder/test-project/service-2/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()
	service_2.RegisterService_2Server(grpcServer, service.NewService2Service(cfg, log, strg))
	reflection.Register(grpcServer)
	return
}

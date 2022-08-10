package services

// import (
import (
	"fmt"

	"github.com/asadbek21coder/test-project/gateway/config"
	"github.com/asadbek21coder/test-project/gateway/genproto/service_1"
	"github.com/asadbek21coder/test-project/gateway/genproto/service_2"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	Service_2() service_2.Service_2Client
	Service_1() service_1.Service_1Client
}

type grpcClients struct {
	service_2 service_2.Service_2Client
	service_1 service_1.Service_1Client
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connService_2, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.Service_2Host, conf.Service_2Port),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	
	connService_1, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.Service_1Host, conf.Service_1Port),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		service_2: service_2.NewService_2Client(connService_2),
		service_1: service_1.NewService_1Client(connService_1),
	}, nil
}

func (g *grpcClients) Service_2() service_2.Service_2Client {
	return g.service_2
}

func (g *grpcClients) Service_1() service_1.Service_1Client {
	return g.service_1
}

package storage

import (
	"context"
	"errors"

	pb "github.com/asadbek21coder/test-project/service-2/genproto/service_2"
)

var ErrorTheSameId = errors.New("cannot use the same uuid for 'id' and 'parent_id' fields")
var ErrorProjectId = errors.New("not valid 'project_id'")

type StorageI interface {
	Service_2_I() Service_2_I
}

type Service_2_I interface {
	GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error)
	GetById(ctx context.Context, req *pb.Id) (*pb.Post, error)
	Update(ctx context.Context, req *pb.Post) (*pb.Post, error)
	Delete(ctx context.Context, req *pb.Id) (*pb.Id, error)
}

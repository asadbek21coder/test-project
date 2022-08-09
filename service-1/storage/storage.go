package storage

import (
	"context"
	"errors"
	pb "service-1/genproto/service_1"
)

var ErrorTheSameId = errors.New("cannot use the same uuid for 'id' and 'parent_id' fields")
var ErrorProjectId = errors.New("not valid 'project_id'")

type StorageI interface {
	Service_1_I() Service_1_I
}

type Service_1_I interface {
	GetAll(ctx context.Context) (*pb.Status, error)
}

package service

import (
	"context"

	"github.com/yukpiz/grpc-e2e-example/external"
	"github.com/yukpiz/grpc-e2e-example/pb"
)

type Service struct {
	External external.ExternalInterface
}

func NewService() *Service {
	return &Service{
		External: external.NewExternal(),
	}
}

func (s *Service) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	s.External.Hello()

	return &pb.HelloResponse{
		Id:   req.Id,
		Name: req.Name,
	}, nil
}

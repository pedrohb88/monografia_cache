package transport

import (
	"monografia/service"
	"monografia/transport/entity"
	pb "monografia/transport/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRouterServer
	service service.Service
	entity  *entity.Entity
}

func NewServer(srv service.Service, entity *entity.Entity) *grpc.Server {

	s := grpc.NewServer(
		grpc.UnaryInterceptor(Benchmark()),
	)

	routerServer := &server{
		service: srv,
		entity:  entity,
	}

	pb.RegisterRouterServer(s, routerServer)

	return s
}

package hello

import (
	"context"

	"github.com/luyasr/hello-world/pkg/ioc"
	"google.golang.org/grpc"
)

var _ Service = (*ServiceImpl)(nil)

type ServiceImpl struct {
	UnimplementedRPCServer
}

func init() {
	ioc.Controller().Registry(&ServiceImpl{})
}

func (s *ServiceImpl) Init() error {
	return nil
}

func (s *ServiceImpl) Name() string {
	return Name
}

func (s *ServiceImpl) Registry(g *grpc.Server) {
	RegisterRPCServer(g, s)
}

func (s *ServiceImpl) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{
		Message: "Hello, " + in.Name,
	}, nil
}

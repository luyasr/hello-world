package hello

import (
	"context"

	"github.com/luyasr/hello-world/pkg/ioc"
)

var _ Service = (*ServiceImpl)(nil)

type ServiceImpl struct {
	UnimplementedRPCServer
}

func init() {
	ioc.Controller().Registry(&ServiceImpl{})
}

func (s *ServiceImpl) Init() error {
	if err := ioc.Controller().Get(Name).Init(); err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) Name() string {
	return Name
}

func (s *ServiceImpl) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{
		Message: "Hello, " + in.Name,
	}, nil
}

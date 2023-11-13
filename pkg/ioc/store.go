package ioc

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Container struct {
	store map[string]Ioc
}

func (c *Container) Init() error {
	for _, obj := range c.store {
		if err := obj.Init(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Container) Get(name string) Ioc {
	return c.store[name]
}

func (c *Container) Registry(obj Ioc) {
	if _, exists := c.store[obj.Name()]; !exists {
		c.store[obj.Name()] = obj
	}
}

// RouteRegistry 断言实现了 GinApiHandler 接口的对象，调用 Registry 方法
func (c *Container) RouteRegistry(r gin.IRouter) {
	for _, route := range c.store {
		if api, ok := route.(GinApiHandler); ok {
			api.Registry(r)
		}
	}
}

// GrpcServiceRegistry 断言实现了 GrpcServerHandler 接口的对象，调用 Registry 方法
func (c *Container) GrpcServiceRegistry(g *grpc.Server) {
	for _, grpc := range c.store {
		if grpcServer, ok := grpc.(GrpcServerHandler); ok {
			grpcServer.Registry(g)
		}
	}
}

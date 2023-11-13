package ioc

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Ioc interface {
	Init() error
	Name() string
}

type GinApiHandler interface {
	Registry(r gin.IRouter)
}

type GrpcServerHandler interface {
	Registry(g *grpc.Server)
}

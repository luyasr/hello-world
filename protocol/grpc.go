package protocol

import (
	"google.golang.org/grpc"
)

type GrpcServer struct {
	grpc *grpc.Server
}

func (g *GrpcServer)Run() {
}
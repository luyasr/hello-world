package protocol

import (
	"log"
	"net"

	"github.com/luyasr/hello-world/config"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	grpc *grpc.Server
}

func NewGrpcServer() *GrpcServer {
	grpcServer := grpc.NewServer()

	return &GrpcServer{
		grpc: grpcServer,
	}
}

func (g *GrpcServer) Run() error {
	listen, err := net.Listen("tcp", config.C.GrpcServer.Addr())
	if err != nil {
		log.Fatal(err)
	}

	log.Print("grpc server running, listen on ", config.C.GrpcServer.Addr())
	return g.grpc.Serve(listen)
}

func (g *GrpcServer) Shutdown() error {
	log.Print("grpc server shutdown")
	g.grpc.GracefulStop()
	
	return nil
}

package ioc

var (
	controller = &Container{
		store: make(map[string]Ioc),
	}
	httpHandler = &Container{
		store: make(map[string]Ioc),
	}
	grpcHandler = &Container{
		store: make(map[string]Ioc),
	}
)

func Controller() *Container {
	return controller
}

func HttpHandler() *Container {
	return httpHandler
}

func GrpcHandler() *Container {
	return grpcHandler
}

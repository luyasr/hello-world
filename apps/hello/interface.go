package hello

const Name = "hello"

type Service interface {
	// 一部分能力通过http暴露
	// 一部分能力通过rpc暴露
	RPCServer
}

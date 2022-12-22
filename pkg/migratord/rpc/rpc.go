package rpc

type RPC interface {
	Run()
}

type rpc struct {
	generated
}

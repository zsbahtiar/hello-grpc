package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type funcServerOption struct {
	f func(*serverOptions)
}

type serverOptions struct {
	grpcPort string
	restPort string
	handlers []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

type ServerOption interface {
	apply(*serverOptions)
}

func newFuncServerOption(f func(*serverOptions)) *funcServerOption {
	return &funcServerOption{
		f: f,
	}
}

func GRPCPort(port string) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.grpcPort = port
	})
}

func RestPort(port string) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.restPort = port
	})
}

func (fdo *funcServerOption) apply(do *serverOptions) {
	fdo.f(do)
}

func RegisterHandler(handler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.handlers = append(o.handlers, handler)
	})
}

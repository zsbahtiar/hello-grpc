package server

type funcServerOption struct {
	f func(*serverOptions)
}

type serverOptions struct {
	grpcPort string
	restPort string
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

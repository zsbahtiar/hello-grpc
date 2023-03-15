package main

import (
	"github.com/zsbahtiar/hello-grpc/auth/core/module"
	"github.com/zsbahtiar/hello-grpc/auth/handler/api"
	pbAuth "github.com/zsbahtiar/hello-grpc/auth/handler/api/grpc/pb/auth"
	"github.com/zsbahtiar/hello-grpc/common/server"
)

func main() {
	authUsecase := module.NewAuthUsecase()
	authHandler := api.NewAuthHandler(authUsecase)

	srv := server.NewServer(
		server.RestPort("8080"),
		server.GRPCPort("9090"),
	)

	pbAuth.RegisterService(srv, authHandler)
	srv.Init()

}

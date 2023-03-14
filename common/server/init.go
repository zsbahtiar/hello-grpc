package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type Server struct {
	Server *grpc.Server
	opts   serverOptions
}

func NewServer(opt ...ServerOption) *Server {
	srv := Server{}
	opts := serverOptions{}
	for _, o := range opt {
		o.apply(&opts)
	}
	srv.Server = grpc.NewServer()
	srv.opts = opts
	return &srv
}

func (s *Server) Init() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.opts.grpcPort))
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}

	go func() {
		log.Printf("starting gRPC server at :%s", s.opts.grpcPort)
		if err := s.Server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
}

func (s *Server) RunGwGracefully() {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("0.0.0.0:%s", s.opts.grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDisableHealthCheck(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	for _, handler := range s.opts.handlers {
		err = handler(ctx, gwmux, conn)
		if err != nil {
			log.Fatal(err)
		}
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.opts.restPort),
		Handler: gwmux,
	}

	log.Printf("starting rest gateway server at :%s", s.opts.restPort)
	if err = gwServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	s.Server.GracefulStop()
	gwServer.Close()

}

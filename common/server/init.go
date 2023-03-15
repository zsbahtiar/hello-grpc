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
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Mux    *runtime.ServeMux
}

func NewServer(opt ...ServerOption) *Server {
	srv := Server{}
	opts := serverOptions{}
	for _, o := range opt {
		o.apply(&opts)
	}
	srv.Server = grpc.NewServer()
	srv.Mux = runtime.NewServeMux()
	srv.Ctx = context.Background()
	srv.opts = opts
	return &srv
}

func (s *Server) Init() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.opts.grpcPort))
	if err != nil {
		log.Fatalf("Error creating listener: %v", err)
	}

	go func() {
		if err := s.Server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	conn, err := grpc.DialContext(
		s.Ctx,
		fmt.Sprintf("0.0.0.0:%s", s.opts.grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDisableHealthCheck(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	s.Conn = conn
	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.opts.restPort),
		Handler: s.Mux,
	}

	log.Printf("starting rest gateway server at :%s", s.opts.restPort)
	log.Printf("starting gRPC server at :%s", s.opts.grpcPort)

	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatalf("failed starting rest gateway: %v", err)
	}

	s.Server.GracefulStop()
	gwServer.Close()
}

package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"yuumi/app/interface/passport/internal/biz/authenticate"
	"yuumi/app/interface/passport/internal/biz/authorize"
	"yuumi/app/interface/passport/internal/biz/github"
	"yuumi/app/interface/passport/internal/conf"
	"yuumi/pkg/logger"
)

type Server struct {
	Addr   string
	server *grpc.Server
}

func (s Server) Start() {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatalf("grpc server failed at: %s \n", s.Addr)
	}

	log.Printf("grpc server listening at: %s \n", s.Addr)
	s.server.Serve(listener)
}

func (s Server) Stop() {
	log.Printf("grpc server shutdown at: %s \n", s.Addr)
	s.server.GracefulStop()
}

func (s *Server) SetServer(opt ...grpc.ServerOption) {
	s.server = grpc.NewServer(opt...)
}

func (s *Server) RegisterServer(log *logger.Logger, c *conf.Config) {
	// ... regist services
	authenticate.RegisterServer(s.server, log, c)
	authorize.RegisterServer(s.server, log, c)
	github.RegisterServer(s.server, log, c)
}

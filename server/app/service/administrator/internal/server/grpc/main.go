package grpc

import (
	"log"
	"net"
	"yuumi/app/service/administrator/internal/biz/administrator"
	"yuumi/app/service/administrator/internal/biz/navmenu"
	"yuumi/app/service/administrator/internal/biz/permission"
	"yuumi/app/service/administrator/internal/biz/role"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
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

func (s *Server) RegisterServer(log *logger.Logger) {
	// ... regist services
	administrator.RegisterServer(s.server, log)
	role.RegisterServer(s.server, log)
	permission.RegisterServer(s.server, log)
	navmenu.RegisterServer(s.server, log)
}

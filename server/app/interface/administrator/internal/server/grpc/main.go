package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"yuumi/app/interface/administrator/internal/biz/administrator"
	"yuumi/app/interface/administrator/internal/biz/navmenu"
	"yuumi/app/interface/administrator/internal/biz/permission"
	"yuumi/app/interface/administrator/internal/biz/role"
	"yuumi/app/interface/administrator/internal/conf"
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

func (s *Server) RegisterServer(log *logger.Logger, serviceconf *conf.ServiceConfig) {
	// ... regist services
	administrator.RegisterServer(s.server, log, serviceconf)
	permission.RegisterServer(s.server, log, serviceconf)
	role.RegisterServer(s.server, log, serviceconf)
	navmenu.RegisterServer(s.server, log, serviceconf)
}

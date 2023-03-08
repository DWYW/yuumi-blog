package grpc

import (
	"log"
	"net"
	"yuumi/app/service/article/internal/biz/article"
	"yuumi/app/service/article/internal/biz/comment"
	"yuumi/app/service/article/internal/biz/visitor"
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
	article.RegisterServer(s.server, log)
	visitor.RegisterServer(s.server, log)
	comment.RegisterServer(s.server, log)
}

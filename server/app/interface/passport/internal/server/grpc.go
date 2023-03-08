package server

import (
	"fmt"
	"yuumi/app/interface/passport/internal/conf"
	"yuumi/pkg/logger"

	grpcServer "yuumi/app/interface/passport/internal/server/grpc"

	"google.golang.org/grpc"
)

func NewGrpcServer(log *logger.Logger, c *conf.Config, opt ...grpc.ServerOption) grpcServer.Server {
	s := grpcServer.Server{
		Addr: fmt.Sprintf("%s:%d", c.Server.Grpc.Host, c.Server.Grpc.Port),
	}
	s.SetServer(opt...)
	s.RegisterServer(log, c)
	return s
}

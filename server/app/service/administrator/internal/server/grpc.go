package server

import (
	"fmt"
	"yuumi/app/service/administrator/internal/conf"
	"yuumi/pkg/logger"

	grpcServer "yuumi/app/service/administrator/internal/server/grpc"

	"google.golang.org/grpc"
)

func NewGrpcServer(log *logger.Logger, opt ...grpc.ServerOption) grpcServer.Server {
	config := conf.GetGRPCServerConfig()
	s := grpcServer.Server{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
	}
	s.SetServer(opt...)
	s.RegisterServer(log)
	return s
}

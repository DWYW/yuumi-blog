package server

import (
	"fmt"

	"yuumi/app/service/article/internal/conf"

	"yuumi/app/service/article/internal/server/getaway"
)

func NewGetawayServer() getaway.Server {
	grpcConfig := conf.GetGRPCServerConfig()
	config := conf.GetHTTPServerConfig()

	s := getaway.Server{
		CorsEnable:          config.Cors.Enabled,
		CorsWhiteListRegexp: config.Cors.WhiteListRegexp,
		Addr:                fmt.Sprintf("%s:%d", config.Host, config.Port),
		GrpcAddr:            fmt.Sprintf("%s:%d", grpcConfig.Host, grpcConfig.Port),
	}
	s.SetMux()
	s.SetContext()
	s.SetServer()
	s.RegisterHandlerFromEndpoint()
	return s
}

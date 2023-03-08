package main

import (
	"flag"
	"yuumi/app/interface/administrator/internal/conf"
	"yuumi/app/interface/administrator/internal/server"
	"yuumi/pkg/interceptor"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

var (
	config = flag.String("config", "app/interface/administrator/config/config.yaml", "config file")
)

func init() {
	flag.Parse()
	initAppConfig()
}

func main() {
	log := createLogger()

	serverGroup := server.ServerGroup{}
	serverGroup.Add(server.NewGrpcServer(
		log,
		conf.GetServiceConfig(),
		grpc.ChainUnaryInterceptor(
			interceptor.GetLogUnaryInterceptor(log),
			interceptor.GetErrorLogUnaryInterceptor(log),
		),
	))
	serverGroup.Add(server.NewHttpServer())
	serverGroup.Start()
}

func initAppConfig() {
	c := conf.Parse(conf.Read(*config))
	conf.Init(c)
}

func createLogger() *logger.Logger {
	logConf := conf.GetLogConfig()
	return logger.New(&logger.Config{
		Dir:          logConf.Dir,
		MaxAgeDay:    int(logConf.MaxAgeDay),
		RotationHour: int(logConf.RotationHour),
	})
}

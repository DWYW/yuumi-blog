package main

import (
	"flag"
	"yuumi/app/service/administrator/internal/conf"
	"yuumi/app/service/administrator/internal/data/mysql"
	"yuumi/app/service/administrator/internal/server"
	"yuumi/pkg/interceptor"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

var (
	config = flag.String("config", "app/service/administrator/config/config.yaml", "config file")
)

func init() {
	flag.Parse()
	initAppConfig()
	initDatabase()
}

func main() {
	log := createLogger()

	serverGroup := server.ServerGroup{}
	serverGroup.Add(server.NewGrpcServer(
		log,
		grpc.ChainUnaryInterceptor(
			interceptor.GetLogUnaryInterceptor(log),
			interceptor.GetErrorLogUnaryInterceptor(log),
			interceptor.GetListUnaryInterceptor(conf.GetPaginationConfig().PageSize),
		),
	))
	serverGroup.Start()
}

func initAppConfig() {
	c := conf.Parse(conf.Read(*config))
	conf.Init(c)
}

func initDatabase() {
	c := conf.GetMysqlConfig()
	db := mysql.NewClient(&mysql.ClientConfig{
		Host:     c.Host,
		Port:     c.Port,
		Username: c.Username,
		Password: c.Password,
		Dbname:   c.Dbname,
		Charset:  c.Charset,
		Pool: &mysql.ClientPoolConfig{
			MaxIdleConns: c.Pool.MaxIdleConns,
			MaxOpenConns: c.Pool.MaxOpenConns,
			MaxLifetime:  c.Pool.MaxLifetime,
		},
	})
	mysql.Init(db)
}

func createLogger() *logger.Logger {
	logConf := conf.GetLogConfig()
	return logger.New(&logger.Config{
		Dir:          logConf.Dir,
		MaxAgeDay:    int(logConf.MaxAgeDay),
		RotationHour: int(logConf.RotationHour),
	})
}

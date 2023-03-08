package authenticate

import (
	v1 "yuumi/api/interface/passport/v1"
	"yuumi/app/interface/passport/internal/conf"
	"yuumi/pkg/client"
	"yuumi/pkg/logger"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 注册服务
func RegisterServer(s *grpc.Server, log *logger.Logger, c *conf.Config) {
	service := Service{
		Log: log,
		AdministratorClient: &client.GrpcClient{
			Host: c.Service.Administrator.Host,
			Port: c.Service.Administrator.Port,
		},
	}

	server := Server{}
	v1.RegisterAuthenticateInterfaceServer(s, &server)

	server.transportHandlers = map[string]grpcTransport.Handler{
		"AdministratorWithJWT": grpcTransport.NewServer(
			MakeAdministratorWithJWTEndpoint(service),
			DecodeAdministratorWithJWTRequest,
			EncodeResponse,
		),
		"VisitorWithJWT": grpcTransport.NewServer(
			MakeVisitorWithJWTEndpoint(service),
			DecodeVisitorWithJWTRequest,
			EncodeResponse,
		),
	}
}

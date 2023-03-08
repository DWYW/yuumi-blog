package permission

import (
	"yuumi/app/interface/administrator/internal/conf"
	"yuumi/middleware"
	"yuumi/pkg/client"
	"yuumi/pkg/logger"

	v1 "yuumi/api/interface/administrator/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 注册服务
func RegisterServer(s *grpc.Server, log *logger.Logger, serviceconf *conf.ServiceConfig) {
	service := Service{
		Log: log,
		AdministratorClient: &client.GrpcClient{
			Host: serviceconf.Administrator.Host,
			Port: serviceconf.Administrator.Port,
		},
	}

	server := Server{Logger: log}
	v1.RegisterPermissionInterfaceServer(s, &server)

	passportHost := serviceconf.Passport.Host
	passportPort := serviceconf.Passport.Port

	server.transportHandlers = map[string]grpcTransport.Handler{
		"Create": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeCreateEndpoint(service),
			),
			DecodeCreateRequest,
			EncodeResponse,
		),
		"Delete": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeDeleteEndpoint(service),
			),
			DecodeDeleteRequest,
			EncodeResponse,
		),
		"Update": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeUpdateEndpoint(service),
			),
			DecodeUpdateRequest,
			EncodeResponse,
		),
		"GetList": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetListEndpoint(service),
			),
			DecodeGetListRequest,
			EncodeResponse,
		),
	}
}

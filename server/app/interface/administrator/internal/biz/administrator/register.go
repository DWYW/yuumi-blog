package administrator

import (
	v1 "yuumi/api/interface/administrator/v1"
	"yuumi/app/interface/administrator/internal/conf"
	"yuumi/middleware"
	"yuumi/pkg/client"
	"yuumi/pkg/logger"

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
	v1.RegisterAdministratorInterfaceServer(s, &server)

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
		"UpdateName": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeUpdateNameEndpoint(service),
			),
			DecodeUpdateNameRequest,
			EncodeResponse,
		),
		"UpdatePassword": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeUpdatePasswordEndpoint(service),
			),
			DecodeUpdatePasswordRequest,
			EncodeResponse,
		),
		"GetInfo": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetInfoEndpoint(service),
			),
			DecodeGetInfoRequest,
			EncodeResponse,
		),
		"GetList": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetListEndpoint(service),
			),
			DecodeGetListRequest,
			EncodeResponse,
		),
		"AppendRolesWithAdministratorID": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeAppendRolesWithAdministratorIDEndpoint(service),
			),
			DecodeAppendRolesWithAdministratorIDRequest,
			EncodeResponse,
		),
		"DeleteRolesWithAdministratorID": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeDeleteRolesWithAdministratorIDEndpoint(service),
			),
			DecodeDeleteRolesWithAdministratorIDRequest,
			EncodeResponse,
		),
	}
}

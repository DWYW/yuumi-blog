package navmenu

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
	v1.RegisterNavMenuInterfaceServer(s, &server)

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
		"GetInfo": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetInfoEndpoint(service),
			),
			DecodeGetInfoRequest,
			EncodeResponse,
		),
		"GetNavMenus": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetNavMenusEndpoint(service),
			),
			DecodeGetNavMenusRequest,
			EncodeResponse,
		),
		"BindWithRoleIDs": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeBindWithRoleIDsEndpoint(service),
			),
			DecodeBindWithRoleIDsRequest,
			EncodeResponse,
		),
		"UnbindWithRoleIDs": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeUnbindWithRoleIDsEndpoint(service),
			),
			DecodeUnbindWithRoleIDsRequest,
			EncodeResponse,
		),
		"GetNavMenusWithAdministratorID": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetNavMenusWithAdministratorIDEndpoint(service),
			),
			DecodeGetNavMenusWithAdministratorIDRequest,
			EncodeResponse,
		),
		"GetNavMenusWithMine": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetNavMenusWithMineEndpoint(service),
			),
			DecodeGetNavMenusWithMineRequest,
			EncodeResponse,
		),
	}
}

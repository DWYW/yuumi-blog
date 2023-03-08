package role

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
	v1.RegisterRoleInterfaceServer(s, &server)

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
		"GetList": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetListEndpoint(service),
			),
			DecodeGetListRequest,
			EncodeResponse,
		),
		"GetRoles": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetRolesEndpoint(service),
			),
			DecodeGetRolesRequest,
			EncodeResponse,
		),
		"GetRolesWithAdministratorID": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeGetRolesWithAdministratorIDEndpoint(service),
			),
			DecodeGetRolesWithAdministratorIDRequest,
			EncodeResponse,
		),
		"AppendPermissionsWithRoleID": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeAppendPermissionsWithRoleIDEndpoint(service),
			),
			DecodeAppendPermissionsWithRoleIDRequest,
			EncodeResponse,
		),
		"DeletePermissionsWithRoleID": grpcTransport.NewServer(
			middleware.AdministratorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeDeletePermissionsWithRoleIDEndpoint(service),
			),
			DecodeDeletePermissionsWithRoleIDRequest,
			EncodeResponse,
		),
	}
}

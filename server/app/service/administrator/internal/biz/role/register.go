package role

import (
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/pkg/logger"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 注册服务
func RegisterServer(s *grpc.Server, log *logger.Logger) {
	service := Service{
		Logger: log,
	}
	server := Server{}
	v1.RegisterRoleServiceServer(s, &server)

	server.transportHandlers = map[string]grpcTransport.Handler{
		"Create": grpcTransport.NewServer(
			MakeCreateEndpoint(service),
			DecodeCreateRequest,
			EncodeResponse,
		),
		"Delete": grpcTransport.NewServer(
			MakeDeleteEndpoint(service),
			DecodeDeleteRequest,
			EncodeResponse,
		),
		"Update": grpcTransport.NewServer(
			MakeUpdateEndpoint(service),
			DecodeUpdateRequest,
			EncodeResponse,
		),
		"GetInfo": grpcTransport.NewServer(
			MakeGetInfoEndpoint(service),
			DecodeGetInfoRequest,
			EncodeResponse,
		),
		"GetInfoWithPermissionID": grpcTransport.NewServer(
			MakeGetInfoWithPermissionIDEndpoint(service),
			DecodeGetInfoWithPermissionIDRequest,
			EncodeResponse,
		),
		"GetList": grpcTransport.NewServer(
			MakeGetListEndpoint(service),
			DecodeGetListRequest,
			EncodeResponse,
		),
		"GetRoles": grpcTransport.NewServer(
			MakeGetRolesEndpoint(service),
			DecodeGetRolesRequest,
			EncodeResponse,
		),
		"GetRolesWithAdministratorID": grpcTransport.NewServer(
			MakeGetRolesWithAdministratorIDEndpoint(service),
			DecodeGetRolesWithAdministratorIDRequest,
			EncodeResponse,
		),
		"AppendPermissionsWithRoleID": grpcTransport.NewServer(
			MakeAppendPermissionsWithRoleIDEndpoint(service),
			DecodeAppendPermissionsWithRoleIDRequest,
			EncodeResponse,
		),
		"DeletePermissionsWithRoleID": grpcTransport.NewServer(
			MakeDeletePermissionsWithRoleIDEndpoint(service),
			DecodeDeletePermissionsWithRoleIDRequest,
			EncodeResponse,
		),
	}
}

package permission

import (
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/app/service/administrator/internal/data/mysql"
	"yuumi/app/service/administrator/internal/data/mysql/permission"
	"yuumi/pkg/logger"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 注册服务
func RegisterServer(s *grpc.Server, log *logger.Logger) {
	service := Service{
		Logger:          log,
		PermissionModel: permission.New(mysql.GetClient()),
	}
	server := Server{}
	v1.RegisterPermissionServiceServer(s, &server)

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
		"GetList": grpcTransport.NewServer(
			MakeGetListEndpoint(service),
			DecodeGetListRequest,
			EncodeResponse,
		),
		"GetPermissions": grpcTransport.NewServer(
			MakeGetPermissionsEndpoint(service),
			DecodeGetPermissionsRequest,
			EncodeResponse,
		),
		"GetPermissionsWithRoleID": grpcTransport.NewServer(
			MakeGetPermissionsWithRoleIDEndpoint(service),
			DecodeGetPermissionsWithRoleIDRequest,
			EncodeResponse,
		),
	}
}

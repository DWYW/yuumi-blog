package navmenu

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
	v1.RegisterNavMenuServiceServer(s, &server)

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
		"GetNavMenus": grpcTransport.NewServer(
			MakeGetNavMenusEndpoint(service),
			DecodeGetNavMenusRequest,
			EncodeResponse,
		),
		"BindWithRoleIDs": grpcTransport.NewServer(
			MakeBindWithRoleIDsEndpoint(service),
			DecodeBindWithRoleIDsRequest,
			EncodeResponse,
		),
		"UnbindWithRoleIDs": grpcTransport.NewServer(
			MakeUnbindWithRoleIDsEndpoint(service),
			DecodeUnbindWithRoleIDsRequest,
			EncodeResponse,
		),
		"GetNavMenusWithAdministratorID": grpcTransport.NewServer(
			MakeGetNavMenusWithAdministratorIDEndpoint(service),
			DecodeGetNavMenusWithAdministratorIDRequest,
			EncodeResponse,
		),
	}
}

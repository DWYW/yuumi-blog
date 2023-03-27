package administrator

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
	v1.RegisterAdministratorServiceServer(s, &server)

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
		"UpdateName": grpcTransport.NewServer(
			MakeUpdateNameEndpoint(service),
			DecodeUpdateNameRequest,
			EncodeResponse,
		),
		"UpdatePassword": grpcTransport.NewServer(
			MakeUpdatePasswordEndpoint(service),
			DecodeUpdatePasswordRequest,
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
		"GetAdministrators": grpcTransport.NewServer(
			MakeGetAdministratorsEndpoint(service),
			DecodeGetAdministratorsRequest,
			EncodeResponse,
		),
		"GetAdministratorWithNameAndPassword": grpcTransport.NewServer(
			MakeGetAdministratorWithNameAndPasswordEndpoint(service),
			DecodeGetAdministratorWithNameAndPasswordRequest,
			EncodeResponse,
		),
		"AppendRolesWithAdministratorID": grpcTransport.NewServer(
			MakeAppendRolesWithAdministratorIDEndpoint(service),
			DecodeAppendRolesWithAdministratorIDRequest,
			EncodeResponse,
		),
		"DeleteRolesWithAdministratorID": grpcTransport.NewServer(
			MakeDeleteRolesWithAdministratorIDEndpoint(service),
			DecodeDeleteRolesWithAdministratorIDRequest,
			EncodeResponse,
		),
	}
}

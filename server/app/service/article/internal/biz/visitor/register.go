package visitor

import (
	v1 "yuumi/api/service/article/v1"
	"yuumi/app/service/article/internal/data/mysql"
	"yuumi/app/service/article/internal/data/mysql/visitor"
	"yuumi/pkg/logger"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 注册服务
func RegisterServer(s *grpc.Server, log *logger.Logger) {
	service := Service{
		Logger:       log,
		VisitorModel: visitor.New(mysql.GetClient()),
	}
	server := Server{}
	v1.RegisterVisitorServiceServer(s, &server)

	server.transportHandlers = map[string]grpcTransport.Handler{
		"CreateWithGithubUser": grpcTransport.NewServer(
			MakeCreateWithGithubUserEndpoint(service),
			DecodeCreateWithGithubUserRequest,
			EncodeResponse,
		),
		"UpdateWithGithubUser": grpcTransport.NewServer(
			MakeUpdateWithGithubUserEndpoint(service),
			DecodeUpdateWithGithubUserRequest,
			EncodeResponse,
		),
		"Delete": grpcTransport.NewServer(
			MakeDeleteEndpoint(service),
			DecodeDeleteRequest,
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
		"GetVisitorWithGithubID": grpcTransport.NewServer(
			MakeGetVisitorWithGithubIDEndpoint(service),
			DecodeGetVisitorWithGithubIDRequest,
			EncodeResponse,
		),
	}
}

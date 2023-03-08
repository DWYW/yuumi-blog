package article

import (
	"yuumi/app/interface/article/internal/conf"
	"yuumi/middleware"
	"yuumi/pkg/client"
	"yuumi/pkg/logger"

	v1 "yuumi/api/interface/article/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 注册服务
func RegisterServer(s *grpc.Server, log *logger.Logger, serviceconf *conf.ServiceConfig) {
	service := Service{
		Log: log,
		ArticleClient: &client.GrpcClient{
			Host: serviceconf.Article.Host,
			Port: serviceconf.Article.Port,
		},
	}

	server := Server{Logger: log}
	v1.RegisterArticleInterfaceServer(s, &server)

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
			MakeGetInfoEndpoint(service),
			DecodeGetInfoRequest,
			EncodeResponse,
		),
		"GetList": grpcTransport.NewServer(
			MakeGetListEndpoint(service),
			DecodeGetListRequest,
			EncodeResponse,
		),
	}
}

package comment

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
	v1.RegisterCommentInterfaceServer(s, &server)

	passportHost := serviceconf.Passport.Host
	passportPort := serviceconf.Passport.Port

	server.transportHandlers = map[string]grpcTransport.Handler{
		"Create": grpcTransport.NewServer(
			middleware.VisitorAuthenticationEndpointMiddleware(passportHost, passportPort)(
				MakeCreateEndpoint(service),
			),
			DecodeCreateRequest,
			EncodeResponse,
		),
		"GetList": grpcTransport.NewServer(
			MakeGetListEndpoint(service),
			DecodeGetListRequest,
			EncodeResponse,
		),
	}
}

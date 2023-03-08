package authorize

import (
	v1 "yuumi/api/interface/passport/v1"
	"yuumi/app/interface/passport/internal/conf"
	"yuumi/pkg/client"
	"yuumi/pkg/logger"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// 注册服务
func RegisterServer(s *grpc.Server, log *logger.Logger, c *conf.Config) {
	service := Service{
		Log: log,
		AdministratorClient: &client.GrpcClient{
			Host: c.Service.Administrator.Host,
			Port: c.Service.Administrator.Port,
		},
		ArticleClient: &client.GrpcClient{
			Host: c.Service.Article.Host,
			Port: c.Service.Article.Port,
		},
	}

	server := Server{}
	v1.RegisterAuthorizeInterfaceServer(s, &server)

	server.transportHandlers = map[string]grpcTransport.Handler{
		"AdministratorWithPassword": grpcTransport.NewServer(
			MakeAdministratorWithPasswordEndpoint(service),
			DecodeAdministratorWithPasswordRequest,
			EncodeResponse,
		),
		"VisitorWithGithubSession": grpcTransport.NewServer(
			MakeVisitorWithGithubSessionEndpoint(service),
			DecodeVisitorWithGithubSessionRequest,
			EncodeResponse,
		),
	}
}

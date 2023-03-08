package github

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
		Log:                log,
		GithubClient:       &client.HttpClient{Host: "", Port: 0},
		GithubClientID:     c.Github.ClientID,
		GithubClientSecret: c.Github.ClientSecret,
	}

	server := Server{}
	v1.RegisterGithubInterfaceServer(s, &server)

	server.transportHandlers = map[string]grpcTransport.Handler{
		"GetSessionWithGithubCode": grpcTransport.NewServer(
			MakeGetSessionWithGithubCodeEndpoint(service),
			DecodeGetSessionWithGithubCodeRequest,
			EncodeResponse,
		),
	}
}

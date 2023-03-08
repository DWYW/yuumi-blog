package github

import (
	"context"
	v1 "yuumi/api/interface/passport/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	v1.UnimplementedGithubInterfaceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) GetSessionWithGithubCode(ctx context.Context, in *v1.GetSessionWithGithubCodeRequest) (*v1.GetSessionWithGithubCodeReply, error) {
	_, rsp, err := s.transportHandlers["GetSessionWithGithubCode"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetSessionWithGithubCodeReply), err
}

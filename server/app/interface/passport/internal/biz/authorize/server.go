package authorize

import (
	"context"
	v1 "yuumi/api/interface/passport/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	v1.UnimplementedAuthorizeInterfaceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) AdministratorWithPassword(ctx context.Context, in *v1.AdministratorWithPasswordAuthorizeRequest) (*v1.AdministratorWithPasswordAuthorizeReply, error) {
	_, rsp, err := s.transportHandlers["AdministratorWithPassword"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.AdministratorWithPasswordAuthorizeReply), err
}

func (s *Server) VisitorWithGithubSession(ctx context.Context, in *v1.VisitorWithGithubSessionAuthorizeRequest) (*v1.VisitorWithGithubSessionAuthorizeReply, error) {
	_, rsp, err := s.transportHandlers["VisitorWithGithubSession"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.VisitorWithGithubSessionAuthorizeReply), err
}

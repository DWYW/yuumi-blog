package authenticate

import (
	"context"
	v1 "yuumi/api/interface/passport/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	v1.UnimplementedAuthenticateInterfaceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) AdministratorWithJWT(ctx context.Context, in *v1.AdministratorWithJWTAuthenticateRequest) (*v1.AdministratorWithJWTAuthenticateReply, error) {
	_, rsp, err := s.transportHandlers["AdministratorWithJWT"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.AdministratorWithJWTAuthenticateReply), err
}

func (s *Server) VisitorWithJWT(ctx context.Context, in *v1.VisitorWithJWTAuthenticateRequest) (*v1.VisitorWithJWTAuthenticateReply, error) {
	_, rsp, err := s.transportHandlers["VisitorWithJWT"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.VisitorWithJWTAuthenticateReply), err
}

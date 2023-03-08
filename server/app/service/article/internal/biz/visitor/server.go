package visitor

import (
	"context"
	v1 "yuumi/api/service/article/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	v1.UnimplementedVisitorServiceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) CreateWithGithubUser(ctx context.Context, in *v1.CreateWithGithubUserRequest) (*v1.CreateWithGithubUserReply, error) {
	_, rsp, err := s.transportHandlers["CreateWithGithubUser"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.CreateWithGithubUserReply), err
}

func (s *Server) UpdateWithGithubUser(ctx context.Context, in *v1.UpdateWithGithubUserRequest) (*v1.UpdateWithGithubUserReply, error) {
	_, rsp, err := s.transportHandlers["UpdateWithGithubUser"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.UpdateWithGithubUserReply), err
}

func (s *Server) Delete(ctx context.Context, in *v1.DeleteVisitorRequest) (*v1.DeleteVisitorReply, error) {
	_, rsp, err := s.transportHandlers["Delete"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.DeleteVisitorReply), err
}

func (s *Server) GetInfo(ctx context.Context, in *v1.GetVisitorInfoRequest) (*v1.GetVisitorInfoReply, error) {
	_, rsp, err := s.transportHandlers["GetInfo"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetVisitorInfoReply), err
}

func (s *Server) GetList(ctx context.Context, in *v1.GetVisitorListRequest) (*v1.GetVisitorListReply, error) {
	_, rsp, err := s.transportHandlers["GetList"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetVisitorListReply), err
}

func (s *Server) GetVisitorWithGithubID(ctx context.Context, in *v1.GetVisitorWithGithubIDRequest) (*v1.GetVisitorWithGithubIDReply, error) {
	_, rsp, err := s.transportHandlers["GetVisitorWithGithubID"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetVisitorWithGithubIDReply), err
}

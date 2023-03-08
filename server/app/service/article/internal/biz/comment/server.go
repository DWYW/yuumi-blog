package comment

import (
	"context"
	v1 "yuumi/api/service/article/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	v1.UnimplementedCommentServiceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) Create(ctx context.Context, in *v1.CreateCommentRequest) (*v1.CreateCommentReply, error) {
	_, rsp, err := s.transportHandlers["Create"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.CreateCommentReply), err
}

func (s *Server) Delete(ctx context.Context, in *v1.DeleteCommentRequest) (*v1.DeleteCommentReply, error) {
	_, rsp, err := s.transportHandlers["Delete"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.DeleteCommentReply), err
}

func (s *Server) GetInfo(ctx context.Context, in *v1.GetCommentInfoRequest) (*v1.GetCommentInfoReply, error) {
	_, rsp, err := s.transportHandlers["GetInfo"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetCommentInfoReply), err
}

func (s *Server) GetList(ctx context.Context, in *v1.GetCommentListRequest) (*v1.GetCommentListReply, error) {
	_, rsp, err := s.transportHandlers["GetList"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetCommentListReply), err
}

package comment

import (
	"context"
	v1 "yuumi/api/interface/article/v1"
	"yuumi/pkg/logger"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	Logger *logger.Logger
	v1.UnimplementedCommentInterfaceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) Create(ctx context.Context, in *v1.CreateCommentRequest) (*v1.CreateCommentReply, error) {
	_, rsp, err := s.transportHandlers["Create"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.CreateCommentReply), err
}

func (s *Server) GetList(ctx context.Context, in *v1.GetCommentListRequest) (*v1.GetCommentListReply, error) {
	_, rsp, err := s.transportHandlers["GetList"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetCommentListReply), err
}

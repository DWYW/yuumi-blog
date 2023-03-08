package permission

import (
	"context"
	v1 "yuumi/api/interface/administrator/v1"
	"yuumi/pkg/logger"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	Logger *logger.Logger
	v1.UnimplementedPermissionInterfaceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) Create(ctx context.Context, in *v1.CreatePermissionRequest) (*v1.CreatePermissionReply, error) {
	_, rsp, err := s.transportHandlers["Create"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.CreatePermissionReply), err
}

func (s *Server) Delete(ctx context.Context, in *v1.DeletePermissionRequest) (*v1.DeletePermissionReply, error) {
	_, rsp, err := s.transportHandlers["Delete"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.DeletePermissionReply), err
}

func (s *Server) Update(ctx context.Context, in *v1.UpdatePermissionRequest) (*v1.UpdatePermissionReply, error) {
	_, rsp, err := s.transportHandlers["Update"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.UpdatePermissionReply), err
}

func (s *Server) GetList(ctx context.Context, in *v1.GetPermissionListRequest) (*v1.GetPermissionListReply, error) {
	_, rsp, err := s.transportHandlers["GetList"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetPermissionListReply), err
}

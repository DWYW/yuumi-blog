package permission

import (
	"context"
	v1 "yuumi/api/service/administrator/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	v1.UnimplementedPermissionServiceServer
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

func (s *Server) GetInfo(ctx context.Context, in *v1.GetPermissionInfoRequest) (*v1.GetPermissionInfoReply, error) {
	_, rsp, err := s.transportHandlers["GetInfo"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetPermissionInfoReply), err
}

func (s *Server) GetList(ctx context.Context, in *v1.GetPermissionListRequest) (*v1.GetPermissionListReply, error) {
	_, rsp, err := s.transportHandlers["GetList"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetPermissionListReply), err
}

func (s *Server) GetPermissions(ctx context.Context, in *v1.GetPermissionsRequest) (*v1.GetPermissionsReply, error) {
	_, rsp, err := s.transportHandlers["GetPermissions"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetPermissionsReply), err
}

func (s *Server) GetPermissionsWithRoleID(ctx context.Context, in *v1.GetPermissionsWithRoleIDRequest) (*v1.GetPermissionsWithRoleIDReply, error) {
	_, rsp, err := s.transportHandlers["GetPermissionsWithRoleID"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetPermissionsWithRoleIDReply), err
}

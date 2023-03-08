package role

import (
	"context"
	v1 "yuumi/api/service/administrator/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	v1.UnimplementedRoleServiceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) Create(ctx context.Context, in *v1.CreateRoleRequest) (*v1.CreateRoleReply, error) {
	_, rsp, err := s.transportHandlers["Create"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.CreateRoleReply), err
}

func (s *Server) Delete(ctx context.Context, in *v1.DeleteRoleRequest) (*v1.DeleteRoleReply, error) {
	_, rsp, err := s.transportHandlers["Delete"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.DeleteRoleReply), err
}

func (s *Server) Update(ctx context.Context, in *v1.UpdateRoleRequest) (*v1.UpdateRoleReply, error) {
	_, rsp, err := s.transportHandlers["Update"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.UpdateRoleReply), err
}

func (s *Server) GetInfo(ctx context.Context, in *v1.GetRoleInfoRequest) (*v1.GetRoleInfoReply, error) {
	_, rsp, err := s.transportHandlers["GetInfo"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetRoleInfoReply), err
}

func (s *Server) GetInfoWithPermissionID(ctx context.Context, in *v1.GetRoleInfoWithPermissionIDRequest) (*v1.GetRoleInfoWithPermissionIDReply, error) {
	_, rsp, err := s.transportHandlers["GetInfoWithPermissionID"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetRoleInfoWithPermissionIDReply), err
}

func (s *Server) GetList(ctx context.Context, in *v1.GetRoleListRequest) (*v1.GetRoleListReply, error) {
	_, rsp, err := s.transportHandlers["GetList"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetRoleListReply), err
}

func (s *Server) GetRoles(ctx context.Context, in *v1.GetRolesRequest) (*v1.GetRolesReply, error) {
	_, rsp, err := s.transportHandlers["GetRoles"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetRolesReply), err
}

func (s *Server) GetRolesWithAdministratorID(ctx context.Context, in *v1.GetRolesWithAdministratorIDRequest) (*v1.GetRolesWithAdministratorIDReply, error) {
	_, rsp, err := s.transportHandlers["GetRolesWithAdministratorID"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetRolesWithAdministratorIDReply), err
}

func (s *Server) AppendPermissionsWithRoleID(ctx context.Context, in *v1.AppendPermissionsWithRoleIDRequest) (*v1.AppendPermissionsWithRoleIDReply, error) {
	_, rsp, err := s.transportHandlers["AppendPermissionsWithRoleID"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.AppendPermissionsWithRoleIDReply), err
}

func (s *Server) DeletePermissionsWithRoleID(ctx context.Context, in *v1.DeletePermissionsWithRoleIDRequest) (*v1.DeletePermissionsWithRoleIDReply, error) {
	_, rsp, err := s.transportHandlers["DeletePermissionsWithRoleID"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.DeletePermissionsWithRoleIDReply), err
}

package administrator

import (
	"context"
	v1 "yuumi/api/service/administrator/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	v1.UnimplementedAdministratorServiceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) Create(ctx context.Context, in *v1.CreateAdministratorRequest) (*v1.CreateAdministratorReply, error) {
	_, rsp, err := s.transportHandlers["Create"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.CreateAdministratorReply), err
}

func (s *Server) Delete(ctx context.Context, in *v1.DeleteAdministratorRequest) (*v1.DeleteAdministratorReply, error) {
	_, rsp, err := s.transportHandlers["Delete"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.DeleteAdministratorReply), err
}

func (s *Server) UpdateName(ctx context.Context, in *v1.UpdateAdministratorNameRequest) (*v1.UpdateAdministratorNameReply, error) {
	_, rsp, err := s.transportHandlers["UpdateName"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.UpdateAdministratorNameReply), err
}

func (s *Server) UpdatePassword(ctx context.Context, in *v1.UpdateAdministratorPasswordRequest) (*v1.UpdateAdministratorPasswordReply, error) {
	_, rsp, err := s.transportHandlers["UpdatePassword"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.UpdateAdministratorPasswordReply), err
}

func (s *Server) GetInfo(ctx context.Context, in *v1.GetAdministratorInfoRequest) (*v1.GetAdministratorInfoReply, error) {
	_, rsp, err := s.transportHandlers["GetInfo"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetAdministratorInfoReply), err
}

func (s *Server) GetList(ctx context.Context, in *v1.GetAdministratorListRequest) (*v1.GetAdministratorListReply, error) {
	_, rsp, err := s.transportHandlers["GetList"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetAdministratorListReply), err
}

func (s *Server) GetAdministrators(ctx context.Context, in *v1.GetAdministratorsRequest) (*v1.GetAdministratorsReply, error) {
	_, rsp, err := s.transportHandlers["GetAdministrators"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetAdministratorsReply), err
}

func (s *Server) GetAdministratorWithNameAndPassword(ctx context.Context, in *v1.GetAdministratorWithNameAndPasswordRequest) (*v1.GetAdministratorWithNameAndPasswordReply, error) {
	_, rsp, err := s.transportHandlers["GetAdministratorWithNameAndPassword"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetAdministratorWithNameAndPasswordReply), err
}

func (s *Server) AppendRolesWithAdministratorID(ctx context.Context, in *v1.AppendRolesWithAdministratorIDRequest) (*v1.AppendRolesWithAdministratorIDReply, error) {
	_, rsp, err := s.transportHandlers["AppendRolesWithAdministratorID"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.AppendRolesWithAdministratorIDReply), err
}

func (s *Server) DeleteRolesWithAdministratorID(ctx context.Context, in *v1.DeleteRolesWithAdministratorIDRequest) (*v1.DeleteRolesWithAdministratorIDReply, error) {
	_, rsp, err := s.transportHandlers["DeleteRolesWithAdministratorID"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.DeleteRolesWithAdministratorIDReply), err
}

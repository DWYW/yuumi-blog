package navmenu

import (
	"context"
	v1 "yuumi/api/service/administrator/v1"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

// 服务实现
type Server struct {
	v1.UnimplementedNavMenuServiceServer
	transportHandlers map[string]grpcTransport.Handler
}

func (s *Server) Create(ctx context.Context, in *v1.CreateNavMenuRequest) (*v1.CreateNavMenuReply, error) {
	_, rsp, err := s.transportHandlers["Create"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.CreateNavMenuReply), err
}

func (s *Server) Delete(ctx context.Context, in *v1.DeleteNavMenuRequest) (*v1.DeleteNavMenuReply, error) {
	_, rsp, err := s.transportHandlers["Delete"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.DeleteNavMenuReply), err
}

func (s *Server) Update(ctx context.Context, in *v1.UpdateNavMenuRequest) (*v1.UpdateNavMenuReply, error) {
	_, rsp, err := s.transportHandlers["Update"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.UpdateNavMenuReply), err
}

func (s *Server) GetInfo(ctx context.Context, in *v1.GetNavMenuInfoRequest) (*v1.GetNavMenuInfoReply, error) {
	_, rsp, err := s.transportHandlers["GetInfo"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetNavMenuInfoReply), err
}

func (s *Server) GetNavMenus(ctx context.Context, in *v1.GetNavMenusRequest) (*v1.GetNavMenusReply, error) {
	_, rsp, err := s.transportHandlers["GetNavMenus"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetNavMenusReply), err
}

func (s *Server) BindWithRoleIDs(ctx context.Context, in *v1.NavMenuBindWithRoleIDsRequest) (*v1.NavMenuBindWithRoleIDsReply, error) {
	_, rsp, err := s.transportHandlers["BindWithRoleIDs"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.NavMenuBindWithRoleIDsReply), err
}

func (s *Server) UnbindWithRoleIDs(ctx context.Context, in *v1.NavMenuUnbindWithRoleIDsRequest) (*v1.NavMenuUnbindWithRoleIDsReply, error) {
	_, rsp, err := s.transportHandlers["UnbindWithRoleIDs"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.NavMenuUnbindWithRoleIDsReply), err
}

func (s *Server) GetNavMenusWithAdministratorID(ctx context.Context, in *v1.GetNavMenusWithAdministratorIDRequest) (*v1.GetNavMenusWithAdministratorIDReply, error) {
	_, rsp, err := s.transportHandlers["GetNavMenusWithAdministratorID"].ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*v1.GetNavMenusWithAdministratorIDReply), err
}

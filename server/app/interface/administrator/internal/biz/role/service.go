package role

import (
	"context"
	v1 "yuumi/api/interface/administrator/v1"
	serviceV1 "yuumi/api/service/administrator/v1"
	"yuumi/pkg/client"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreateRoleRequest) (*v1.CreateRoleReply, error)
	Delete(ctx context.Context, in *v1.DeleteRoleRequest) (*v1.DeleteRoleReply, error)
	Update(ctx context.Context, in *v1.UpdateRoleRequest) (*v1.UpdateRoleReply, error)
	GetInfo(ctx context.Context, in *v1.GetRoleInfoRequest) (*v1.GetRoleInfoReply, error)
	GetList(ctx context.Context, in *v1.GetRoleListRequest) (*v1.GetRoleListReply, error)
	GetRoles(ctx context.Context, in *v1.GetRolesRequest) (*v1.GetRolesReply, error)
	GetRolesWithAdministratorID(ctx context.Context, in *v1.GetRolesWithAdministratorIDRequest) (*v1.GetRolesWithAdministratorIDReply, error)
	AppendPermissionsWithRoleID(ctx context.Context, in *v1.AppendPermissionsWithRoleIDRequest) (*v1.AppendPermissionsWithRoleIDReply, error)
	DeletePermissionsWithRoleID(ctx context.Context, in *v1.DeletePermissionsWithRoleIDRequest) (*v1.DeletePermissionsWithRoleIDReply, error)
}

type Service struct {
	Log                 *logger.Logger
	AdministratorClient *client.GrpcClient
}

func (s Service) Create(ctx context.Context, in *v1.CreateRoleRequest) (*v1.CreateRoleReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.CreateRoleReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewRoleServiceClient(cc)
		reply, err = client.Create(ctx, &serviceV1.CreateRoleRequest{
			Name:     in.Name,
			Type:     in.Type,
			ParentId: in.ParentId,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.CreateRoleReply{
		Role: reply.Role,
	}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteRoleRequest) (*v1.DeleteRoleReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.DeleteRoleReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewRoleServiceClient(cc)
		reply, err = client.Delete(ctx, &serviceV1.DeleteRoleRequest{Id: in.Id})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.DeleteRoleReply{
		Message: reply.Message,
	}, nil
}

func (s Service) Update(ctx context.Context, in *v1.UpdateRoleRequest) (*v1.UpdateRoleReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.UpdateRoleReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewRoleServiceClient(cc)
		reply, err = client.Update(ctx, &serviceV1.UpdateRoleRequest{
			Id:       in.Id,
			Name:     in.Name,
			Type:     in.Type,
			ParentId: in.ParentId,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.UpdateRoleReply{
		Role: reply.Role,
	}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetRoleInfoRequest) (*v1.GetRoleInfoReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetRoleInfoReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewRoleServiceClient(cc)
		reply, err = client.GetInfo(ctx, &serviceV1.GetRoleInfoRequest{Id: in.Id, PreloadPermissions: in.PreloadPermissions})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetRoleInfoReply{
		Role: reply.Role,
	}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetRoleListRequest) (*v1.GetRoleListReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetRoleListReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewRoleServiceClient(cc)
		reply, err = client.GetList(ctx, &serviceV1.GetRoleListRequest{Page: in.Page, PageSize: in.PageSize})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetRoleListReply{
		Roles:      reply.Roles,
		Pagination: reply.Pagination,
	}, nil
}

func (s Service) GetRoles(ctx context.Context, in *v1.GetRolesRequest) (*v1.GetRolesReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetRolesReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewRoleServiceClient(cc)
		reply, err = client.GetRoles(ctx, &serviceV1.GetRolesRequest{})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetRolesReply{
		Roles: reply.Roles,
	}, nil
}

func (s Service) GetRolesWithAdministratorID(ctx context.Context, in *v1.GetRolesWithAdministratorIDRequest) (*v1.GetRolesWithAdministratorIDReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetRolesWithAdministratorIDReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewRoleServiceClient(cc)
		reply, err = client.GetRolesWithAdministratorID(ctx, &serviceV1.GetRolesWithAdministratorIDRequest{AdministratorId: in.AdministratorId})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetRolesWithAdministratorIDReply{Roles: reply.Roles}, nil
}

func (s Service) AppendPermissionsWithRoleID(ctx context.Context, in *v1.AppendPermissionsWithRoleIDRequest) (*v1.AppendPermissionsWithRoleIDReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.AppendPermissionsWithRoleIDReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewRoleServiceClient(cc)
		reply, err = client.AppendPermissionsWithRoleID(ctx, &serviceV1.AppendPermissionsWithRoleIDRequest{RoleId: in.Id, PermissionIds: in.PermissionIds})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.AppendPermissionsWithRoleIDReply{
		Message: reply.Message,
	}, nil
}

func (s Service) DeletePermissionsWithRoleID(ctx context.Context, in *v1.DeletePermissionsWithRoleIDRequest) (*v1.DeletePermissionsWithRoleIDReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.DeletePermissionsWithRoleIDReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewRoleServiceClient(cc)
		reply, err = client.DeletePermissionsWithRoleID(ctx, &serviceV1.DeletePermissionsWithRoleIDRequest{RoleId: in.Id, PermissionIds: in.PermissionIds})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.DeletePermissionsWithRoleIDReply{
		Message: reply.Message,
	}, nil
}

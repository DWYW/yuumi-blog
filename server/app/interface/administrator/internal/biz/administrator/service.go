package administrator

import (
	"context"
	v1 "yuumi/api/interface/administrator/v1"
	serviceV1 "yuumi/api/service/administrator/v1"

	"yuumi/pkg/client"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreateAdministratorRequest) (*v1.CreateAdministratorReply, error)
	Delete(ctx context.Context, in *v1.DeleteAdministratorRequest) (*v1.DeleteAdministratorReply, error)
	UpdateName(ctx context.Context, in *v1.UpdateAdministratorNameRequest) (*v1.UpdateAdministratorNameReply, error)
	UpdatePassword(ctx context.Context, in *v1.UpdateAdministratorPasswordRequest) (*v1.UpdateAdministratorPasswordReply, error)
	GetInfo(ctx context.Context, in *v1.GetAdministratorInfoRequest) (*v1.GetAdministratorInfoReply, error)
	GetList(ctx context.Context, in *v1.GetAdministratorListRequest) (*v1.GetAdministratorListReply, error)
	AppendRolesWithAdministratorID(ctx context.Context, in *v1.AppendRolesWithAdministratorIDRequest) (*v1.AppendRolesWithAdministratorIDReply, error)
	DeleteRolesWithAdministratorID(ctx context.Context, in *v1.DeleteRolesWithAdministratorIDRequest) (*v1.DeleteRolesWithAdministratorIDReply, error)
}

type Service struct {
	Log                 *logger.Logger
	AdministratorClient *client.GrpcClient
}

func (s Service) Create(ctx context.Context, in *v1.CreateAdministratorRequest) (*v1.CreateAdministratorReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.CreateAdministratorReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewAdministratorServiceClient(cc)
		reply, err = client.Create(ctx, &serviceV1.CreateAdministratorRequest{
			Name:     in.Name,
			Password: in.Password,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.CreateAdministratorReply{
		Administrator: reply.Administrator,
	}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteAdministratorRequest) (*v1.DeleteAdministratorReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.DeleteAdministratorReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewAdministratorServiceClient(cc)
		reply, err = client.Delete(ctx, &serviceV1.DeleteAdministratorRequest{Id: in.Id})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.DeleteAdministratorReply{
		Message: reply.Message,
	}, nil
}

func (s Service) UpdateName(ctx context.Context, in *v1.UpdateAdministratorNameRequest) (*v1.UpdateAdministratorNameReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.UpdateAdministratorNameReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewAdministratorServiceClient(cc)
		reply, err = client.UpdateName(ctx, &serviceV1.UpdateAdministratorNameRequest{Id: in.Id, Name: in.Name})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.UpdateAdministratorNameReply{
		Administrator: reply.Administrator,
	}, nil
}

func (s Service) UpdatePassword(ctx context.Context, in *v1.UpdateAdministratorPasswordRequest) (*v1.UpdateAdministratorPasswordReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.UpdateAdministratorPasswordReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewAdministratorServiceClient(cc)
		reply, err = client.UpdatePassword(ctx, &serviceV1.UpdateAdministratorPasswordRequest{Id: in.Id, Password: in.Password, PasswordNew: in.PasswordNew})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.UpdateAdministratorPasswordReply{
		Administrator: reply.Administrator,
	}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetAdministratorInfoRequest) (*v1.GetAdministratorInfoReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetAdministratorInfoReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewAdministratorServiceClient(cc)
		reply, err = client.GetInfo(ctx, &serviceV1.GetAdministratorInfoRequest{Id: in.Id, PreloadRoles: in.PreloadRoles})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetAdministratorInfoReply{
		Administrator: reply.Administrator,
	}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetAdministratorListRequest) (*v1.GetAdministratorListReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetAdministratorListReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewAdministratorServiceClient(cc)
		reply, err = client.GetList(ctx, &serviceV1.GetAdministratorListRequest{
			Page:         in.Page,
			PageSize:     in.PageSize,
			PreloadRoles: in.PreloadRoles,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetAdministratorListReply{
		Administrators: reply.Administrators,
		Pagination:     reply.Pagination,
	}, nil
}

func (s Service) AppendRolesWithAdministratorID(ctx context.Context, in *v1.AppendRolesWithAdministratorIDRequest) (*v1.AppendRolesWithAdministratorIDReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.AppendRolesWithAdministratorIDReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewAdministratorServiceClient(cc)
		reply, err = client.AppendRolesWithAdministratorID(ctx, &serviceV1.AppendRolesWithAdministratorIDRequest{AdministratorId: in.Id, RoleIds: in.RoleIds})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.AppendRolesWithAdministratorIDReply{
		Message: reply.Message,
	}, nil
}

func (s Service) DeleteRolesWithAdministratorID(ctx context.Context, in *v1.DeleteRolesWithAdministratorIDRequest) (*v1.DeleteRolesWithAdministratorIDReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.DeleteRolesWithAdministratorIDReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewAdministratorServiceClient(cc)
		reply, err = client.DeleteRolesWithAdministratorID(ctx, &serviceV1.DeleteRolesWithAdministratorIDRequest{AdministratorId: in.Id, RoleIds: in.RoleIds})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.DeleteRolesWithAdministratorIDReply{
		Message: reply.Message,
	}, nil
}

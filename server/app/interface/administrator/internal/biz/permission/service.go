package permission

import (
	"context"
	v1 "yuumi/api/interface/administrator/v1"
	administratorV1 "yuumi/api/service/administrator/v1"
	"yuumi/pkg/client"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreatePermissionRequest) (*v1.CreatePermissionReply, error)
	Delete(ctx context.Context, in *v1.DeletePermissionRequest) (*v1.DeletePermissionReply, error)
	Update(ctx context.Context, in *v1.UpdatePermissionRequest) (*v1.UpdatePermissionReply, error)
	GetList(ctx context.Context, in *v1.GetPermissionListRequest) (*v1.GetPermissionListReply, error)
}

type Service struct {
	Log                 *logger.Logger
	AdministratorClient *client.GrpcClient
}

func (s Service) Create(ctx context.Context, in *v1.CreatePermissionRequest) (*v1.CreatePermissionReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *administratorV1.CreatePermissionReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := administratorV1.NewPermissionServiceClient(cc)
		reply, err = client.Create(ctx, &administratorV1.CreatePermissionRequest{Name: in.Name, RpcMethod: in.RpcMethod})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.CreatePermissionReply{
		Permission: reply.Permission,
	}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeletePermissionRequest) (*v1.DeletePermissionReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *administratorV1.DeletePermissionReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := administratorV1.NewPermissionServiceClient(cc)
		reply, err = client.Delete(ctx, &administratorV1.DeletePermissionRequest{Id: in.Id})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.DeletePermissionReply{
		Message: reply.Message,
	}, nil
}

func (s Service) Update(ctx context.Context, in *v1.UpdatePermissionRequest) (*v1.UpdatePermissionReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *administratorV1.UpdatePermissionReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := administratorV1.NewPermissionServiceClient(cc)
		reply, err = client.Update(ctx, &administratorV1.UpdatePermissionRequest{
			Id:        in.Id,
			Name:      in.Name,
			RpcMethod: in.RpcMethod,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.UpdatePermissionReply{
		Permission: reply.Permission,
	}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetPermissionListRequest) (*v1.GetPermissionListReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *administratorV1.GetPermissionListReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := administratorV1.NewPermissionServiceClient(cc)
		reply, err = client.GetList(ctx, &administratorV1.GetPermissionListRequest{
			Page:         in.Page,
			PageSize:     in.PageSize,
			PreloadRoles: in.PreloadRoles,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetPermissionListReply{
		Permissions: reply.Permissions,
		Pagination:  reply.Pagination,
	}, nil
}

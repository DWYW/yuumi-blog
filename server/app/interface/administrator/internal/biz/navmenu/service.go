package navmenu

import (
	"context"
	v1 "yuumi/api/interface/administrator/v1"
	serviceV1 "yuumi/api/service/administrator/v1"
	"yuumi/pkg/client"
	"yuumi/pkg/keys"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreateNavMenuRequest) (*v1.CreateNavMenuReply, error)
	Delete(ctx context.Context, in *v1.DeleteNavMenuRequest) (*v1.DeleteNavMenuReply, error)
	Update(ctx context.Context, in *v1.UpdateNavMenuRequest) (*v1.UpdateNavMenuReply, error)
	GetInfo(ctx context.Context, in *v1.GetNavMenuInfoRequest) (*v1.GetNavMenuInfoReply, error)
	GetNavMenus(ctx context.Context, in *v1.GetNavMenusRequest) (*v1.GetNavMenusReply, error)
	BindWithRoleIDs(ctx context.Context, in *v1.NavMenuBindWithRoleIDsRequest) (*v1.NavMenuBindWithRoleIDsReply, error)
	UnbindWithRoleIDs(ctx context.Context, in *v1.NavMenuUnbindWithRoleIDsRequest) (*v1.NavMenuUnbindWithRoleIDsReply, error)
	GetNavMenusWithAdministratorID(ctx context.Context, in *v1.GetNavMenusWithAdministratorIDRequest) (*v1.GetNavMenusWithAdministratorIDReply, error)
	GetNavMenusWithMine(ctx context.Context, in *v1.GetNavMenusWithMineRequest) (*v1.GetNavMenusWithMineReply, error)
}

type Service struct {
	Log                 *logger.Logger
	AdministratorClient *client.GrpcClient
}

func (s Service) Create(ctx context.Context, in *v1.CreateNavMenuRequest) (*v1.CreateNavMenuReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.CreateNavMenuReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewNavMenuServiceClient(cc)
		reply, err = client.Create(ctx, &serviceV1.CreateNavMenuRequest{
			ParentId:    in.ParentId,
			Name:        in.Name,
			LinkUrl:     in.LinkUrl,
			ActivedRule: in.ActivedRule,
			Weight:      in.Weight,
			Icon:        in.Icon,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.CreateNavMenuReply{Menu: reply.Menu}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteNavMenuRequest) (*v1.DeleteNavMenuReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.DeleteNavMenuReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewNavMenuServiceClient(cc)
		reply, err = client.Delete(ctx, &serviceV1.DeleteNavMenuRequest{Id: in.Id})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.DeleteNavMenuReply{Message: reply.Message}, nil
}

func (s Service) Update(ctx context.Context, in *v1.UpdateNavMenuRequest) (*v1.UpdateNavMenuReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.UpdateNavMenuReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewNavMenuServiceClient(cc)
		reply, err = client.Update(ctx, &serviceV1.UpdateNavMenuRequest{
			Id:          in.Id,
			Name:        in.Name,
			ParentId:    in.ParentId,
			LinkUrl:     in.LinkUrl,
			ActivedRule: in.ActivedRule,
			Weight:      in.Weight,
			Icon:        in.Icon,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.UpdateNavMenuReply{Menu: reply.Menu}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetNavMenuInfoRequest) (*v1.GetNavMenuInfoReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetNavMenuInfoReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewNavMenuServiceClient(cc)
		reply, err = client.GetInfo(ctx, &serviceV1.GetNavMenuInfoRequest{Id: in.Id})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetNavMenuInfoReply{Menu: reply.Menu}, nil
}

func (s Service) GetNavMenus(ctx context.Context, in *v1.GetNavMenusRequest) (*v1.GetNavMenusReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetNavMenusReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewNavMenuServiceClient(cc)
		reply, err = client.GetNavMenus(ctx, &serviceV1.GetNavMenusRequest{
			PreloadRoles: in.PreloadRoles,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return &v1.GetNavMenusReply{Menus: reply.Menus}, nil
}

func (s Service) BindWithRoleIDs(ctx context.Context, in *v1.NavMenuBindWithRoleIDsRequest) (*v1.NavMenuBindWithRoleIDsReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewNavMenuServiceClient(cc)
		_, err = client.BindWithRoleIDs(ctx, &serviceV1.NavMenuBindWithRoleIDsRequest{
			Id:      in.Id,
			RoleIds: in.RoleIds,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.NavMenuBindWithRoleIDsReply{}, nil
}

func (s Service) UnbindWithRoleIDs(ctx context.Context, in *v1.NavMenuUnbindWithRoleIDsRequest) (*v1.NavMenuUnbindWithRoleIDsReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewNavMenuServiceClient(cc)
		_, err = client.UnbindWithRoleIDs(ctx, &serviceV1.NavMenuUnbindWithRoleIDsRequest{
			Id:      in.Id,
			RoleIds: in.RoleIds,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.NavMenuUnbindWithRoleIDsReply{}, nil
}

func (s Service) GetNavMenusWithAdministratorID(ctx context.Context, in *v1.GetNavMenusWithAdministratorIDRequest) (*v1.GetNavMenusWithAdministratorIDReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetNavMenusWithAdministratorIDReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewNavMenuServiceClient(cc)
		reply, err = client.GetNavMenusWithAdministratorID(ctx, &serviceV1.GetNavMenusWithAdministratorIDRequest{AdministratorId: in.Id})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetNavMenusWithAdministratorIDReply{Menus: reply.Menus}, nil
}

func (s Service) GetNavMenusWithMine(ctx context.Context, in *v1.GetNavMenusWithMineRequest) (*v1.GetNavMenusWithMineReply, error) {
	info := keys.GetAdministratorAuthDataFromContext(ctx)

	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *serviceV1.GetNavMenusWithAdministratorIDReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := serviceV1.NewNavMenuServiceClient(cc)
		reply, err = client.GetNavMenusWithAdministratorID(ctx, &serviceV1.GetNavMenusWithAdministratorIDRequest{AdministratorId: info.ID})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetNavMenusWithMineReply{Menus: reply.Menus}, nil
}

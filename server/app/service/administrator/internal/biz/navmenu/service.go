package navmenu

import (
	"context"
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/app/service/administrator/internal/data/mysql/navmenu"
	"yuumi/app/service/administrator/internal/data/mysql/role"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/app/service/administrator/internal/data/mysql/transform"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/logger"
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
}

type Service struct {
	Logger       *logger.Logger
	NavMenuModel *navmenu.Model
	RoleModel    *role.Model
}

func (s Service) Create(ctx context.Context, in *v1.CreateNavMenuRequest) (*v1.CreateNavMenuReply, error) {
	v := schema.NavMenu{Name: in.Name}
	if in.LinkUrl != nil {
		v.LinkUrl = *in.LinkUrl
	}
	if in.ActivedRule != nil {
		v.ActivedRule = *in.ActivedRule
	}
	if in.Weight != nil {
		v.Weight = *in.Weight
	}
	if in.Icon != nil {
		v.Icon = *in.Icon
	}
	if in.ParentId != nil {
		v.ParentID = uint(*in.ParentId)
	}

	res, err := s.NavMenuModel.CreateOne(ctx, &v)
	if err != nil {
		return nil, errorcode.New(errorcode.RecordCreationFailed)
	}

	return &v1.CreateNavMenuReply{Menu: transform.TransformNavMenuData(res)}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteNavMenuRequest) (*v1.DeleteNavMenuReply, error) {
	_, err := s.NavMenuModel.DeleteWithID(ctx, in.Id)
	if err != nil {
		return nil, errorcode.New(errorcode.RecordDeletionFailed)
	}

	return &v1.DeleteNavMenuReply{Message: "ok"}, nil
}

func (s Service) Update(ctx context.Context, in *v1.UpdateNavMenuRequest) (*v1.UpdateNavMenuReply, error) {
	fields := schema.NavMenu{}.GetFeilds()
	target := map[string]interface{}{}
	if in.Name != "" {
		target[fields.Name] = in.Name
	}
	if in.ParentId != nil {
		target[fields.ParentID] = in.ParentId
	}
	if in.LinkUrl != nil {
		target[fields.LinkUrl] = in.LinkUrl
	}
	if in.ActivedRule != nil {
		target[fields.ActivedRule] = in.ActivedRule
	}
	if in.Icon != nil {
		target[fields.Icon] = in.Icon
	}
	if in.Weight != nil {
		target[fields.Weight] = in.Weight
	}

	res, err := s.NavMenuModel.UpdateWithID(ctx, in.Id, target)
	if err != nil {
		return nil, errorcode.New(errorcode.RecordUpdateFailed)
	}

	return &v1.UpdateNavMenuReply{Menu: transform.TransformNavMenuData(res)}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetNavMenuInfoRequest) (*v1.GetNavMenuInfoReply, error) {
	res, err := s.NavMenuModel.FindOne(ctx, &navmenu.FindCondition{
		ID:   in.Id,
		Name: in.Name,
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordNotFound, err)
	}

	return &v1.GetNavMenuInfoReply{Menu: transform.TransformNavMenuData(res)}, nil
}

func (s Service) GetNavMenus(ctx context.Context, in *v1.GetNavMenusRequest) (*v1.GetNavMenusReply, error) {
	res, err := s.NavMenuModel.Find(ctx, &navmenu.FindCondition{
		Sort: &navmenu.FindConditionSort{
			Weight: -1,
		},
		Preload: &navmenu.FindConditionPreload{
			Roles: *in.PreloadRoles,
		},
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetNavMenusReply{Menus: transform.TransformNavMenusData(res)}, nil
}

func (s Service) BindWithRoleIDs(ctx context.Context, in *v1.NavMenuBindWithRoleIDsRequest) (*v1.NavMenuBindWithRoleIDsReply, error) {
	err := s.NavMenuModel.BindWithRoleIDs(ctx, in.Id, in.RoleIds)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RolesBindFailed, err)
	}
	return &v1.NavMenuBindWithRoleIDsReply{}, nil
}

func (s Service) UnbindWithRoleIDs(ctx context.Context, in *v1.NavMenuUnbindWithRoleIDsRequest) (*v1.NavMenuUnbindWithRoleIDsReply, error) {
	err := s.NavMenuModel.UnbindWithRoleIDs(ctx, in.Id, in.RoleIds)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RolesUnbindFailed, err)
	}
	return &v1.NavMenuUnbindWithRoleIDsReply{}, nil
}

func (s Service) GetNavMenusWithAdministratorID(ctx context.Context, in *v1.GetNavMenusWithAdministratorIDRequest) (*v1.GetNavMenusWithAdministratorIDReply, error) {
	// 查找与administrator直接关联的role
	roles, err := s.RoleModel.FindWithAdministrator(ctx, &role.FindWithAdministratorCondition{AdministratorID: in.AdministratorId})
	if err != nil {
		return nil, err
	}

	var roleIDs []int64
	for _, item := range roles {
		roleIDs = append(roleIDs, int64(item.ID))
	}

	// 查找role的children role
	if len(roleIDs) > 0 {
		children, err := s.RoleModel.FindWithParentIds(ctx, roleIDs...)
		if err != nil {
			return nil, err
		}

		for _, item := range children {
			roleIDs = append(roleIDs, int64(item.ID))
		}
	}

	// 查找navmenus
	menus, err := s.NavMenuModel.FindWithRoleIDs(ctx, &navmenu.FindWithRoleIDsCondition{RoleIDS: roleIDs})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetNavMenusWithAdministratorIDReply{Menus: transform.TransformNavMenusData(menus)}, nil
}

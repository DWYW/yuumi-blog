package role

import (
	"context"
	"math"
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/app/service/administrator/internal/data/mysql"
	"yuumi/app/service/administrator/internal/data/mysql/role"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/app/service/administrator/internal/data/mysql/transform"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/logger"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreateRoleRequest) (*v1.CreateRoleReply, error)
	Delete(ctx context.Context, in *v1.DeleteRoleRequest) (*v1.DeleteRoleReply, error)
	Update(ctx context.Context, in *v1.UpdateRoleRequest) (*v1.UpdateRoleReply, error)
	GetInfo(ctx context.Context, in *v1.GetRoleInfoRequest) (*v1.GetRoleInfoReply, error)
	GetInfoWithPermissionID(ctx context.Context, in *v1.GetRoleInfoWithPermissionIDRequest) (*v1.GetRoleInfoWithPermissionIDReply, error)
	GetList(ctx context.Context, in *v1.GetRoleListRequest) (*v1.GetRoleListReply, error)
	GetRoles(ctx context.Context, in *v1.GetRolesRequest) (*v1.GetRolesReply, error)
	GetRolesWithAdministratorID(ctx context.Context, in *v1.GetRolesWithAdministratorIDRequest) (*v1.GetRolesWithAdministratorIDReply, error)
	AppendPermissionsWithRoleID(ctx context.Context, in *v1.AppendPermissionsWithRoleIDRequest) (*v1.AppendPermissionsWithRoleIDReply, error)
	DeletePermissionsWithRoleID(ctx context.Context, in *v1.DeletePermissionsWithRoleIDRequest) (*v1.DeletePermissionsWithRoleIDReply, error)
}

type Service struct {
	Logger *logger.Logger
}

func (s Service) Create(ctx context.Context, in *v1.CreateRoleRequest) (*v1.CreateRoleReply, error) {
	// 通过Name查询是否存在
	count, err := mysql.GetRole().Count(ctx, &role.FindConditionWhere{Type: in.Type}, true)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.Unknown, err)
	} else if count > 0 {
		return nil, errorcode.New(errorcode.RoleTypeExisted)
	}

	// 插入文档
	res, err := mysql.GetRole().CreateOne(ctx, &schema.Role{
		Name:     in.Name,
		Type:     in.Type,
		ParentID: uint(in.ParentId),
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCreationFailed, err)
	}

	return &v1.CreateRoleReply{
		Role: transform.TransformRoleData(res),
	}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteRoleRequest) (*v1.DeleteRoleReply, error) {
	_, err := mysql.GetRole().DeleteWithID(ctx, in.Id)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordDeletionFailed, err)
	}
	return &v1.DeleteRoleReply{Message: "ok"}, nil
}

func (s Service) Update(ctx context.Context, in *v1.UpdateRoleRequest) (*v1.UpdateRoleReply, error) {
	res, err := mysql.GetRole().UpdateWithID(ctx, in.Id, &schema.Role{
		Name:     in.Name,
		Type:     in.Type,
		ParentID: uint(in.ParentId),
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordUpdateFailed, err)
	}

	return &v1.UpdateRoleReply{
		Role: transform.TransformRoleData(res),
	}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetRoleInfoRequest) (*v1.GetRoleInfoReply, error) {
	res, err := mysql.GetRole().FindOne(ctx, &role.FindCondition{
		Where: &role.FindConditionWhere{
			ID:   in.Id,
			Type: in.Type,
		},
		Preload: &role.FindConditionPreload{
			Permissions: in.PreloadPermissions,
		},
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordNotFound, err)
	}

	return &v1.GetRoleInfoReply{
		Role: transform.TransformRoleData(res),
	}, nil
}

func (s Service) GetInfoWithPermissionID(ctx context.Context, in *v1.GetRoleInfoWithPermissionIDRequest) (*v1.GetRoleInfoWithPermissionIDReply, error) {
	res, err := mysql.GetRole().FindWithPermission(ctx, &role.FindWithPermissionCondition{PermissionID: in.PermissionId})
	if err != nil || len(res) == 0 {
		return nil, errorcode.NewWithDetail(errorcode.RecordNotFound, err)
	}

	return &v1.GetRoleInfoWithPermissionIDReply{
		Role: transform.TransformRoleData(res[0]),
	}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetRoleListRequest) (*v1.GetRoleListReply, error) {
	findCondition := &role.FindCondition{
		Where: &role.FindConditionWhere{},
	}

	total, err := mysql.GetRole().Count(ctx, findCondition.Where, true)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCountFailed, err)
	}

	result, err := mysql.GetRole().FindList(ctx, &role.FindListCondition{Page: in.Page, PageSize: in.PageSize, FindCondition: findCondition})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetRoleListReply{
		Roles: transform.TransformRolesData(result),
		Pagination: &v1.PaginationData{
			Page:      in.Page,
			PageSize:  in.PageSize,
			Total:     total,
			PageTotal: int64(math.Ceil(float64(total) / float64(in.PageSize))),
		},
	}, nil
}

func (s Service) GetRoles(ctx context.Context, in *v1.GetRolesRequest) (*v1.GetRolesReply, error) {
	result, err := mysql.GetRole().Find(ctx, &role.FindCondition{})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetRolesReply{
		Roles: transform.TransformRolesData(result),
	}, nil
}

func (s Service) GetRolesWithAdministratorID(ctx context.Context, in *v1.GetRolesWithAdministratorIDRequest) (*v1.GetRolesWithAdministratorIDReply, error) {
	roles, err := mysql.GetRole().FindWithAdministrator(ctx, &role.FindWithAdministratorCondition{AdministratorID: in.AdministratorId})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	// 获取被继承者
	var ids = make([]int64, len(roles))
	for i, role := range roles {
		ids[i] = int64(role.ID)
	}
	res, err := mysql.GetRole().FindWithParentIds(ctx, ids...)
	if err == nil {
		roles = append(roles, res...)
	}

	return &v1.GetRolesWithAdministratorIDReply{
		Roles: transform.TransformRolesData(roles),
	}, nil
}

func (s Service) AppendPermissionsWithRoleID(ctx context.Context, in *v1.AppendPermissionsWithRoleIDRequest) (*v1.AppendPermissionsWithRoleIDReply, error) {
	err := mysql.GetRole().AppendPermissionsWithRoleID(ctx, in.RoleId, in.PermissionIds)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.PermissionsBindFailed, err)
	}

	return &v1.AppendPermissionsWithRoleIDReply{
		Message: "ok",
	}, nil
}

func (s Service) DeletePermissionsWithRoleID(ctx context.Context, in *v1.DeletePermissionsWithRoleIDRequest) (*v1.DeletePermissionsWithRoleIDReply, error) {
	err := mysql.GetRole().DeletePermissionsWithRoleID(ctx, in.RoleId, in.PermissionIds)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.PermissionsUnbindFailed, err)
	}

	return &v1.DeletePermissionsWithRoleIDReply{
		Message: "ok",
	}, nil
}

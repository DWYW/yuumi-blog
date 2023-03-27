package permission

import (
	"context"
	"math"
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/app/service/administrator/internal/data/mysql"
	"yuumi/app/service/administrator/internal/data/mysql/permission"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/app/service/administrator/internal/data/mysql/transform"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/logger"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreatePermissionRequest) (*v1.CreatePermissionReply, error)
	Delete(ctx context.Context, in *v1.DeletePermissionRequest) (*v1.DeletePermissionReply, error)
	Update(ctx context.Context, in *v1.UpdatePermissionRequest) (*v1.UpdatePermissionReply, error)
	GetInfo(ctx context.Context, in *v1.GetPermissionInfoRequest) (*v1.GetPermissionInfoReply, error)
	GetList(ctx context.Context, in *v1.GetPermissionListRequest) (*v1.GetPermissionListReply, error)
	GetPermissions(ctx context.Context, in *v1.GetPermissionsRequest) (*v1.GetPermissionsReply, error)
	GetPermissionsWithRoleID(ctx context.Context, in *v1.GetPermissionsWithRoleIDRequest) (*v1.GetPermissionsWithRoleIDReply, error)
}

type Service struct {
	Logger *logger.Logger
}

func (s Service) Create(ctx context.Context, in *v1.CreatePermissionRequest) (*v1.CreatePermissionReply, error) {
	res, err := mysql.GetPermission().CreateOne(ctx, &schema.Permission{Name: in.Name, RpcMethod: in.RpcMethod})
	if err != nil {
		return nil, errorcode.New(errorcode.RecordCreationFailed)
	}

	return &v1.CreatePermissionReply{
		Permission: transform.TransformPermissionData(res),
	}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeletePermissionRequest) (*v1.DeletePermissionReply, error) {
	_, err := mysql.GetPermission().DeleteWithID(ctx, in.Id)
	if err != nil {
		return nil, errorcode.New(errorcode.RecordDeletionFailed)
	}
	return &v1.DeletePermissionReply{Message: "ok"}, nil
}

func (s Service) Update(ctx context.Context, in *v1.UpdatePermissionRequest) (*v1.UpdatePermissionReply, error) {
	fields := schema.Permission{}.GetFields()
	target := map[string]interface{}{}
	if in.Name != nil {
		target[fields.Name] = *in.Name
	}
	if in.RpcMethod != nil {
		target[fields.RpcMethod] = *in.RpcMethod
	}

	res, err := mysql.GetPermission().UpdateWithID(ctx, in.Id, target)
	if err != nil {
		return nil, errorcode.New(errorcode.RecordDeletionFailed)
	}
	return &v1.UpdatePermissionReply{
		Permission: transform.TransformPermissionData(res),
	}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetPermissionInfoRequest) (*v1.GetPermissionInfoReply, error) {
	res, err := mysql.GetPermission().FindOne(ctx, &permission.FindCondition{
		Where: &permission.FindConditionWhere{ID: in.Id, RpcMethod: in.RpcMethod},
	})
	if err != nil {
		return nil, errorcode.New(errorcode.RecordNotFound)
	}
	return &v1.GetPermissionInfoReply{
		Permission: transform.TransformPermissionData(res),
	}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetPermissionListRequest) (*v1.GetPermissionListReply, error) {
	findCondition := &permission.FindCondition{
		Where: &permission.FindConditionWhere{},
		Preload: &permission.FindConditionPreload{
			Roles: *in.PreloadRoles,
		},
	}

	total, err := mysql.GetPermission().Count(ctx, findCondition.Where, true)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCountFailed, err)
	}

	res, err := mysql.GetPermission().FindList(ctx, &permission.FindListCondition{
		Page:          in.Page,
		PageSize:      in.PageSize,
		FindCondition: findCondition,
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetPermissionListReply{
		Permissions: transform.TransformPermissionsData(res),
		Pagination: &v1.PaginationData{
			Page:      in.Page,
			PageSize:  in.PageSize,
			Total:     total,
			PageTotal: int64(math.Ceil(float64(total) / float64(in.PageSize))),
		},
	}, nil
}

func (s Service) GetPermissions(ctx context.Context, in *v1.GetPermissionsRequest) (*v1.GetPermissionsReply, error) {
	res, err := mysql.GetPermission().Find(ctx, &permission.FindCondition{})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}
	return &v1.GetPermissionsReply{
		Permissions: transform.TransformPermissionsData(res),
	}, nil
}

func (s Service) GetPermissionsWithRoleID(ctx context.Context, in *v1.GetPermissionsWithRoleIDRequest) (*v1.GetPermissionsWithRoleIDReply, error) {
	res, err := mysql.GetPermission().FindWithRole(ctx, &permission.FindWithRoleCondition{RoleID: in.RoleId})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetPermissionsWithRoleIDReply{
		Permissions: transform.TransformPermissionsData(res),
	}, nil
}

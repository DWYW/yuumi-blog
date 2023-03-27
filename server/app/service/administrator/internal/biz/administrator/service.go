package administrator

import (
	"context"
	"math"
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/app/service/administrator/internal/data/mysql"
	"yuumi/app/service/administrator/internal/data/mysql/administrator"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/app/service/administrator/internal/data/mysql/transform"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/logger"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreateAdministratorRequest) (*v1.CreateAdministratorReply, error)
	Delete(ctx context.Context, in *v1.DeleteAdministratorRequest) (*v1.DeleteAdministratorReply, error)
	UpdateName(ctx context.Context, in *v1.UpdateAdministratorNameRequest) (*v1.UpdateAdministratorNameReply, error)
	UpdatePassword(ctx context.Context, in *v1.UpdateAdministratorPasswordRequest) (*v1.UpdateAdministratorPasswordReply, error)
	GetInfo(ctx context.Context, in *v1.GetAdministratorInfoRequest) (*v1.GetAdministratorInfoReply, error)
	GetList(ctx context.Context, in *v1.GetAdministratorListRequest) (*v1.GetAdministratorListReply, error)
	GetAdministrators(ctx context.Context, in *v1.GetAdministratorsRequest) (*v1.GetAdministratorsReply, error)
	GetAdministratorWithNameAndPassword(ctx context.Context, in *v1.GetAdministratorWithNameAndPasswordRequest) (*v1.GetAdministratorWithNameAndPasswordReply, error)
	AppendRolesWithAdministratorID(ctx context.Context, in *v1.AppendRolesWithAdministratorIDRequest) (*v1.AppendRolesWithAdministratorIDReply, error)
	DeleteRolesWithAdministratorID(ctx context.Context, in *v1.DeleteRolesWithAdministratorIDRequest) (*v1.DeleteRolesWithAdministratorIDReply, error)
}

type Service struct {
	Logger *logger.Logger
}

func (s Service) administratorIsExistedWithName(ctx context.Context, name string) error {
	count, err := mysql.GetAdministrator().Count(ctx, &administrator.FindConditionWhere{Name: name}, true)
	if err != nil {
		return errorcode.NewWithDetail(errorcode.Unknown, err)
	} else if count > 0 {
		return errorcode.New(errorcode.AdministratorExisted)
	}

	return nil
}

func (s Service) Create(ctx context.Context, in *v1.CreateAdministratorRequest) (*v1.CreateAdministratorReply, error) {
	if err := s.administratorIsExistedWithName(ctx, in.Name); err != nil {
		return nil, err
	}

	// 插入文档
	res, err := mysql.GetAdministrator().CreateOne(ctx, &schema.Administrator{
		Name:     in.Name,
		Password: in.Password,
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCreationFailed, err)
	}

	return &v1.CreateAdministratorReply{
		Administrator: transform.TransformAdministratorData(res),
	}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteAdministratorRequest) (*v1.DeleteAdministratorReply, error) {
	_, err := mysql.GetAdministrator().DeleteWithID(ctx, in.Id)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordDeletionFailed, err)
	}

	return &v1.DeleteAdministratorReply{Message: "ok"}, nil
}

func (s Service) UpdateName(ctx context.Context, in *v1.UpdateAdministratorNameRequest) (*v1.UpdateAdministratorNameReply, error) {
	if err := s.administratorIsExistedWithName(ctx, in.Name); err != nil {
		return nil, err
	}

	res, err := mysql.GetAdministrator().UpdateWithID(ctx, in.Id, &schema.Administrator{Name: in.Name})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordUpdateFailed, err)
	}

	return &v1.UpdateAdministratorNameReply{
		Administrator: transform.TransformAdministratorData(res),
	}, nil
}

func (s Service) UpdatePassword(ctx context.Context, in *v1.UpdateAdministratorPasswordRequest) (*v1.UpdateAdministratorPasswordReply, error) {
	// 查找administrator信息
	dest, err := mysql.GetAdministrator().FindOne(ctx, &administrator.FindCondition{
		Where: &administrator.FindConditionWhere{ID: in.Id},
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordNotFound, err)
	}

	// 比对密码
	if !mysql.GetAdministrator().PasswordValidator(in.Password, dest.Password, dest.Salt) {
		return nil, errorcode.New(errorcode.PasswordError)
	}

	// 重置密码
	res, err := mysql.GetAdministrator().UpdatePassword(ctx, in.Id, in.PasswordNew, dest.Salt)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordUpdateFailed, err)
	}

	return &v1.UpdateAdministratorPasswordReply{
		Administrator: transform.TransformAdministratorData(res),
	}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetAdministratorInfoRequest) (*v1.GetAdministratorInfoReply, error) {
	res, err := mysql.GetAdministrator().FindOne(ctx, &administrator.FindCondition{
		Where:   &administrator.FindConditionWhere{ID: in.Id},
		Preload: &administrator.FindConditionPreload{Roles: in.PreloadRoles},
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordNotFound, err)
	}

	return &v1.GetAdministratorInfoReply{
		Administrator: transform.TransformAdministratorData(res),
	}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetAdministratorListRequest) (*v1.GetAdministratorListReply, error) {
	findCondition := &administrator.FindCondition{
		Where:   &administrator.FindConditionWhere{},
		Preload: &administrator.FindConditionPreload{Roles: in.PreloadRoles},
	}

	total, err := mysql.GetAdministrator().Count(ctx, findCondition.Where, true)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCountFailed, err)
	}

	result, err := mysql.GetAdministrator().FindList(ctx, &administrator.FindListCondition{
		Page:          in.Page,
		PageSize:      in.PageSize,
		FindCondition: findCondition,
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetAdministratorListReply{
		Administrators: transform.TransformAdministratorsData(result),
		Pagination: &v1.PaginationData{
			Page:      in.Page,
			PageSize:  in.PageSize,
			Total:     total,
			PageTotal: int64(math.Ceil(float64(total) / float64(in.PageSize))),
		},
	}, nil
}

func (s Service) GetAdministrators(ctx context.Context, in *v1.GetAdministratorsRequest) (*v1.GetAdministratorsReply, error) {
	result, err := mysql.GetAdministrator().Find(ctx, &administrator.FindCondition{
		Preload: &administrator.FindConditionPreload{Roles: in.PreloadRoles},
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetAdministratorsReply{
		Administrators: transform.TransformAdministratorsData(result),
	}, nil
}

func (s Service) GetAdministratorWithNameAndPassword(ctx context.Context, in *v1.GetAdministratorWithNameAndPasswordRequest) (*v1.GetAdministratorWithNameAndPasswordReply, error) {
	res, err := mysql.GetAdministrator().FindOne(ctx, &administrator.FindCondition{
		Where:   &administrator.FindConditionWhere{Name: in.Name},
		Preload: &administrator.FindConditionPreload{Roles: true},
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.NameAndPasswordError, err)
	}

	if in.Password != "" && !mysql.GetAdministrator().PasswordValidator(in.Password, res.Password, res.Salt) {
		return nil, errorcode.New(errorcode.NameAndPasswordError)
	}

	return &v1.GetAdministratorWithNameAndPasswordReply{
		Administrator: transform.TransformAdministratorData(res),
	}, nil
}

func (s Service) AppendRolesWithAdministratorID(ctx context.Context, in *v1.AppendRolesWithAdministratorIDRequest) (*v1.AppendRolesWithAdministratorIDReply, error) {
	err := mysql.GetAdministrator().AppendRolesWithAdministratorID(ctx, in.AdministratorId, in.RoleIds)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RolesBindFailed, err)
	}

	return &v1.AppendRolesWithAdministratorIDReply{Message: "ok"}, nil
}

func (s Service) DeleteRolesWithAdministratorID(ctx context.Context, in *v1.DeleteRolesWithAdministratorIDRequest) (*v1.DeleteRolesWithAdministratorIDReply, error) {
	err := mysql.GetAdministrator().DeleteRolesWithAdministratorID(ctx, in.AdministratorId, in.RoleIds)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RolesUnbindFailed, err)
	}

	return &v1.DeleteRolesWithAdministratorIDReply{Message: "ok"}, nil
}

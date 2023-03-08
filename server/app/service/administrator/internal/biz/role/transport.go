package role

import (
	"context"
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/keys"
)

func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

func DecodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.CreateRoleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Name == "" || request.Type == "" {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeleteRoleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.IDRequired)
	}

	return request, nil
}

func DecodeUpdateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.UpdateRoleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.IDRequired)
	}

	return request, nil
}

func DecodeGetInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetRoleInfoRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	return request, nil
}

func DecodeGetInfoWithPermissionIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetRoleInfoWithPermissionIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.PermissionId == 0 {
		return nil, errorcode.New(errorcode.MethodRequired)
	}

	return request, nil
}

func DecodeGetListRequest(ctx context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetRoleListRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Page == 0 {
		request.Page = 1
	}

	pageSize := keys.GetListPageMaxNumFromContext(ctx)
	if request.PageSize <= 0 || request.PageSize > pageSize {
		request.PageSize = pageSize
	}

	return request, nil
}

func DecodeGetRolesRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetRolesRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	return request, nil
}

func DecodeGetRolesWithAdministratorIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetRolesWithAdministratorIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.AdministratorId < 0 {
		return nil, errorcode.New(errorcode.AdministratorIDRequired)
	}

	return request, nil
}

func DecodeAppendPermissionsWithRoleIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.AppendPermissionsWithRoleIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.RoleId == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeDeletePermissionsWithRoleIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeletePermissionsWithRoleIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.RoleId == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

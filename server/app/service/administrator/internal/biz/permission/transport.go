package permission

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
	request, ok := r.(*v1.CreatePermissionRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Name == "" || request.RpcMethod == "" {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeletePermissionRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.IDRequired)
	}

	return request, nil
}

func DecodeUpdateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.UpdatePermissionRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.IDRequired)
	}

	return request, nil
}

func DecodeGetInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetPermissionInfoRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 && request.RpcMethod == "" {
		return nil, errorcode.New(errorcode.IDRequired)
	}

	return request, nil
}

func DecodeGetListRequest(ctx context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetPermissionListRequest)
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

func DecodeGetPermissionsRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetPermissionsRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}

func DecodeGetPermissionsWithRoleIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetPermissionsWithRoleIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.RoleId == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

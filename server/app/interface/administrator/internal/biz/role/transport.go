package role

import (
	"context"
	v1 "yuumi/api/interface/administrator/v1"
	"yuumi/pkg/errorcode"
)

func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

func DecodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.CreateRoleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}

func DecodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeleteRoleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}
func DecodeUpdateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.UpdateRoleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
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

func DecodeGetListRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetRoleListRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
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
	return request, nil
}

func DecodeAppendPermissionsWithRoleIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.AppendPermissionsWithRoleIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}

func DecodeDeletePermissionsWithRoleIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeletePermissionsWithRoleIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}

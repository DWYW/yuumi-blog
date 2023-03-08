package navmenu

import (
	"context"
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/pkg/errorcode"
)

func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

func DecodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.CreateNavMenuRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Name == "" {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeleteNavMenuRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeUpdateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.UpdateNavMenuRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeGetInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetNavMenuInfoRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 && request.Name == "" {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeGetNavMenusRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetNavMenusRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	return request, nil
}

func DecodeBindWithRoleIDsRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.NavMenuBindWithRoleIDsRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 || len(request.RoleIds) == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeUnbindWithRoleIDsRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.NavMenuUnbindWithRoleIDsRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 || len(request.RoleIds) == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeGetNavMenusWithAdministratorIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetNavMenusWithAdministratorIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.AdministratorId == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

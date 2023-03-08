package administrator

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
	request, ok := r.(*v1.CreateAdministratorRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Name == "" || request.Password == "" {
		return nil, errorcode.New(errorcode.NameAndPasswordRequired)
	}

	return request, nil
}

func DecodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeleteAdministratorRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.IDRequired)
	}

	return request, nil
}

func DecodeUpdateNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.UpdateAdministratorNameRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.IDRequired)
	}

	if request.Name == "" {
		return nil, errorcode.New(errorcode.NameRequired)
	}

	return request, nil
}

func DecodeUpdatePasswordRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.UpdateAdministratorPasswordRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.IDRequired)
	}

	if request.Password == "" || request.PasswordNew == "" {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeGetInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetAdministratorInfoRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	return request, nil
}

func DecodeGetListRequest(ctx context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetAdministratorListRequest)
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

func DecodeGetAdministratorsRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetAdministratorsRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	return request, nil
}

func DecodeGetAdministratorWithNameAndPasswordRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetAdministratorWithNameAndPasswordRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Name == "" || request.Password == "" {
		return nil, errorcode.New(errorcode.NameAndPasswordRequired)
	}

	return request, nil
}

func DecodeAppendRolesWithAdministratorIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.AppendRolesWithAdministratorIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.AdministratorId == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeDeleteRolesWithAdministratorIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeleteRolesWithAdministratorIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.AdministratorId == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

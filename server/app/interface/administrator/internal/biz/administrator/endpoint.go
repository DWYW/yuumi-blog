package administrator

import (
	"context"
	v1 "yuumi/api/interface/administrator/v1"

	"github.com/go-kit/kit/endpoint"
)

func MakeCreateEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.CreateAdministratorRequest)
		return s.Create(ctx, r)
	}
}

func MakeDeleteEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.DeleteAdministratorRequest)
		return s.Delete(ctx, r)
	}
}

func MakeUpdateNameEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.UpdateAdministratorNameRequest)
		return s.UpdateName(ctx, r)
	}
}

func MakeUpdatePasswordEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.UpdateAdministratorPasswordRequest)
		return s.UpdatePassword(ctx, r)
	}
}

func MakeGetInfoEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetAdministratorInfoRequest)
		return s.GetInfo(ctx, r)
	}
}

func MakeGetListEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetAdministratorListRequest)
		return s.GetList(ctx, r)
	}
}

func MakeAppendRolesWithAdministratorIDEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.AppendRolesWithAdministratorIDRequest)
		return s.AppendRolesWithAdministratorID(ctx, r)
	}
}

func MakeDeleteRolesWithAdministratorIDEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.DeleteRolesWithAdministratorIDRequest)
		return s.DeleteRolesWithAdministratorID(ctx, r)
	}
}

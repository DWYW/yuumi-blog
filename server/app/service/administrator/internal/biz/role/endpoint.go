package role

import (
	"context"
	v1 "yuumi/api/service/administrator/v1"

	"github.com/go-kit/kit/endpoint"
)

func MakeCreateEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.CreateRoleRequest)
		return s.Create(ctx, r)
	}
}

func MakeDeleteEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.DeleteRoleRequest)
		return s.Delete(ctx, r)
	}
}

func MakeUpdateEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.UpdateRoleRequest)
		return s.Update(ctx, r)
	}
}

func MakeGetInfoEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetRoleInfoRequest)
		return s.GetInfo(ctx, r)
	}
}

func MakeGetInfoWithPermissionIDEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetRoleInfoWithPermissionIDRequest)
		return s.GetInfoWithPermissionID(ctx, r)
	}
}

func MakeGetListEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetRoleListRequest)
		return s.GetList(ctx, r)
	}
}

func MakeGetRolesEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetRolesRequest)
		return s.GetRoles(ctx, r)
	}
}

func MakeGetRolesWithAdministratorIDEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetRolesWithAdministratorIDRequest)
		return s.GetRolesWithAdministratorID(ctx, r)
	}
}

func MakeAppendPermissionsWithRoleIDEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.AppendPermissionsWithRoleIDRequest)
		return s.AppendPermissionsWithRoleID(ctx, r)
	}
}

func MakeDeletePermissionsWithRoleIDEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.DeletePermissionsWithRoleIDRequest)
		return s.DeletePermissionsWithRoleID(ctx, r)
	}
}

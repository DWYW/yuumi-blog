package permission

import (
	"context"
	v1 "yuumi/api/service/administrator/v1"

	"github.com/go-kit/kit/endpoint"
)

func MakeCreateEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.CreatePermissionRequest)
		return s.Create(ctx, r)
	}
}

func MakeDeleteEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.DeletePermissionRequest)
		return s.Delete(ctx, r)
	}
}

func MakeUpdateEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.UpdatePermissionRequest)
		return s.Update(ctx, r)
	}
}

func MakeGetInfoEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetPermissionInfoRequest)
		return s.GetInfo(ctx, r)
	}
}

func MakeGetListEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetPermissionListRequest)
		return s.GetList(ctx, r)
	}
}

func MakeGetPermissionsEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetPermissionsRequest)
		return s.GetPermissions(ctx, r)
	}
}

func MakeGetPermissionsWithRoleIDEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetPermissionsWithRoleIDRequest)
		return s.GetPermissionsWithRoleID(ctx, r)
	}
}

package permission

import (
	"context"
	v1 "yuumi/api/interface/administrator/v1"

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

func MakeGetListEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetPermissionListRequest)
		return s.GetList(ctx, r)
	}
}

package navmenu

import (
	"context"
	v1 "yuumi/api/interface/administrator/v1"

	"github.com/go-kit/kit/endpoint"
)

func MakeCreateEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.CreateNavMenuRequest)
		return s.Create(ctx, r)
	}
}

func MakeDeleteEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.DeleteNavMenuRequest)
		return s.Delete(ctx, r)
	}
}

func MakeUpdateEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.UpdateNavMenuRequest)
		return s.Update(ctx, r)
	}
}

func MakeGetInfoEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetNavMenuInfoRequest)
		return s.GetInfo(ctx, r)
	}
}

func MakeGetNavMenusEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetNavMenusRequest)
		return s.GetNavMenus(ctx, r)
	}
}

func MakeBindWithRoleIDsEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.NavMenuBindWithRoleIDsRequest)
		return s.BindWithRoleIDs(ctx, r)
	}
}

func MakeUnbindWithRoleIDsEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.NavMenuUnbindWithRoleIDsRequest)
		return s.UnbindWithRoleIDs(ctx, r)
	}
}

func MakeGetNavMenusWithAdministratorIDEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetNavMenusWithAdministratorIDRequest)
		return s.GetNavMenusWithAdministratorID(ctx, r)
	}
}

func MakeGetNavMenusWithMineEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetNavMenusWithMineRequest)
		return s.GetNavMenusWithMine(ctx, r)
	}
}

package visitor

import (
	"context"
	v1 "yuumi/api/service/article/v1"

	"github.com/go-kit/kit/endpoint"
)

func MakeCreateWithGithubUserEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.CreateWithGithubUserRequest)
		return s.CreateWithGithubUser(ctx, r)
	}
}

func MakeUpdateWithGithubUserEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.UpdateWithGithubUserRequest)
		return s.UpdateWithGithubUser(ctx, r)
	}
}

func MakeDeleteEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.DeleteVisitorRequest)
		return s.Delete(ctx, r)
	}
}

func MakeGetInfoEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetVisitorInfoRequest)
		return s.GetInfo(ctx, r)
	}
}

func MakeGetListEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetVisitorListRequest)
		return s.GetList(ctx, r)
	}
}

func MakeGetVisitorWithGithubIDEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetVisitorWithGithubIDRequest)
		return s.GetVisitorWithGithubID(ctx, r)
	}
}

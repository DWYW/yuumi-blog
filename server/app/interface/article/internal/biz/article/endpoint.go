package article

import (
	"context"
	v1 "yuumi/api/interface/article/v1"

	"github.com/go-kit/kit/endpoint"
)

func MakeCreateEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.CreateArticleRequest)
		return s.Create(ctx, r)
	}
}

func MakeDeleteEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.DeleteArticleRequest)
		return s.Delete(ctx, r)
	}
}

func MakeUpdateEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.UpdateArticleRequest)
		return s.Update(ctx, r)
	}
}

func MakeGetInfoEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetArticleInfoRequest)
		return s.GetInfo(ctx, r)
	}
}

func MakeGetListEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetArticleListRequest)
		return s.GetList(ctx, r)
	}
}

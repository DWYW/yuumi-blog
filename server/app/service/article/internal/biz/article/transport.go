package article

import (
	"context"
	v1 "yuumi/api/service/article/v1"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/keys"
)

func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

func DecodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.CreateArticleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	return request, nil
}

func DecodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeleteArticleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeUpdateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.UpdateArticleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeGetInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetArticleInfoRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == nil {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeGetListRequest(ctx context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetArticleListRequest)
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

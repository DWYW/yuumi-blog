package article

import (
	"context"
	v1 "yuumi/api/interface/article/v1"
	"yuumi/pkg/errorcode"
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
	return request, nil
}

func DecodeUpdateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.UpdateArticleRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}

func DecodeGetInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetArticleInfoRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}

func DecodeGetListRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetArticleListRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}

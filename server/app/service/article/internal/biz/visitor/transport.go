package visitor

import (
	"context"
	v1 "yuumi/api/service/article/v1"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/keys"
)

func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

func DecodeCreateWithGithubUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.CreateWithGithubUserRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.GithubUser == nil {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeUpdateWithGithubUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.UpdateWithGithubUserRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 || request.GithubUser == nil {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.DeleteVisitorRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Id == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeGetInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetVisitorInfoRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if *request.Id == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

func DecodeGetListRequest(ctx context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetVisitorListRequest)
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

func DecodeGetVisitorWithGithubIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetVisitorWithGithubIDRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.GithubId == 0 {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

package github

import (
	"context"
	v1 "yuumi/api/interface/passport/v1"
	"yuumi/pkg/errorcode"
)

func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

func DecodeGetSessionWithGithubCodeRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetSessionWithGithubCodeRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Code == "" {
		return nil, errorcode.New(errorcode.MissingRequiredParams)
	}

	return request, nil
}

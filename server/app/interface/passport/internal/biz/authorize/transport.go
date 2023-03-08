package authorize

import (
	"context"
	v1 "yuumi/api/interface/passport/v1"
	"yuumi/pkg/errorcode"
)

func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

func DecodeAdministratorWithPasswordRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.AdministratorWithPasswordAuthorizeRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Name == "" || request.Password == "" {
		return nil, errorcode.New(errorcode.NameAndPasswordRequired)
	}

	return request, nil
}

func DecodeVisitorWithGithubSessionRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.VisitorWithGithubSessionAuthorizeRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	if request.Session == "" {
		return nil, errorcode.New(errorcode.NameAndPasswordRequired)
	}

	return request, nil
}

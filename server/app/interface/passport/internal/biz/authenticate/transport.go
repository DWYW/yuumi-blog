package authenticate

import (
	"context"
	v1 "yuumi/api/interface/passport/v1"
	"yuumi/pkg/errorcode"
)

func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

func DecodeAdministratorWithJWTRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.AdministratorWithJWTAuthenticateRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	return request, nil
}

func DecodeVisitorWithJWTRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.VisitorWithJWTAuthenticateRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}

	return request, nil
}

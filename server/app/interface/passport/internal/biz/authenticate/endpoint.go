package authenticate

import (
	"context"
	v1 "yuumi/api/interface/passport/v1"

	"github.com/go-kit/kit/endpoint"
)

func MakeAdministratorWithJWTEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.AdministratorWithJWTAuthenticateRequest)
		return s.AdministratorWithJWT(ctx, r)
	}
}

func MakeVisitorWithJWTEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.VisitorWithJWTAuthenticateRequest)
		return s.VisitorWithJWT(ctx, r)
	}
}

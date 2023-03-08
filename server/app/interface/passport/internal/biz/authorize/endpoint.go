package authorize

import (
	"context"
	v1 "yuumi/api/interface/passport/v1"

	"github.com/go-kit/kit/endpoint"
)

func MakeAdministratorWithPasswordEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.AdministratorWithPasswordAuthorizeRequest)
		return s.AdministratorWithPassword(ctx, r)
	}
}

func MakeVisitorWithGithubSessionEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.VisitorWithGithubSessionAuthorizeRequest)
		return s.VisitorWithGithubSession(ctx, r)
	}
}

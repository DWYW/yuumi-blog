package github

import (
	"context"
	v1 "yuumi/api/interface/passport/v1"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetSessionWithGithubCodeEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, _ := request.(*v1.GetSessionWithGithubCodeRequest)
		return s.GetSessionWithGithubCode(ctx, r)
	}
}

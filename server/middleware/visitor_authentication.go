package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	v1 "yuumi/api/interface/passport/v1"
	"yuumi/pkg/keys"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func VisitorAuthenticationEndpointMiddleware(host string, port uint64) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			conn, err := grpc.Dial(
				fmt.Sprintf("%s:%d", host, port),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithPerRPCCredentials(JWTAuthToken{
					Token:     keys.GetAuthorizationFromContext(ctx),
					RequestID: keys.GetRequestIDFormContext(ctx),
				}),
			)

			if err != nil {
				bytes, _ := json.Marshal(struct {
					ErrCode int64  `json:"err_code"`
					ErrMsg  string `json:"err_msg"`
				}{ErrCode: 500, ErrMsg: "authentication service can not connect"})

				return nil, errors.New(string(bytes))
			}

			defer conn.Close()

			client := v1.NewAuthenticateInterfaceClient(conn)
			reply, err := client.VisitorWithJWT(ctx, &v1.VisitorWithJWTAuthenticateRequest{Method: keys.GetRequestFullMethodFromContext(ctx)})
			if err != nil {
				return nil, err
			}

			// 添加到上下文中
			ctx = keys.SetVisitorAuthData(ctx, &keys.VisitorAuthData{ID: reply.Id, Name: reply.Name})
			return next(ctx, request)
		}
	}
}

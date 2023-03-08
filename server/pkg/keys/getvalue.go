package keys

import (
	"context"
	"time"

	"google.golang.org/grpc/metadata"
)

func GetAuthorizationFromContext(ctx context.Context) string {
	var token string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md["authorization"]) > 0 {
		token = md["authorization"][0]
	}

	return token
}

func GetRequestIDFormContext(ctx context.Context) string {
	v, _ := ctx.Value(RequestIDKey).(string)
	return v
}

func GetRequestFullMethodFromContext(ctx context.Context) string {
	v, _ := ctx.Value(RequestFullMethodKey).(string)
	return v
}

func GetRequestStartTimeFromContext(ctx context.Context) time.Time {
	v, _ := ctx.Value(RequestStartTimeKey).(time.Time)
	return v
}

func GetListPageMaxNumFromContext(ctx context.Context) int64 {
	v, _ := ctx.Value(ListPageMaxNumKey).(int64)
	return v
}

func GetAdministratorAuthDataFromContext(ctx context.Context) *AdministratorAuthData {
	v, _ := ctx.Value(AdministratorAuthDataKey).(*AdministratorAuthData)
	return v
}

func GetVisitorAuthDataFromContext(ctx context.Context) *VisitorAuthData {
	v, _ := ctx.Value(VisitorAuthDataKey).(*VisitorAuthData)
	return v
}

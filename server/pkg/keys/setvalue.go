package keys

import (
	"context"
	"time"
)

func SetRequestID(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, RequestIDKey, v)
}

func SetRequestFullMethod(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, RequestFullMethodKey, v)
}

func SetRequestStartTime(ctx context.Context, v time.Time) context.Context {
	return context.WithValue(ctx, RequestStartTimeKey, v)
}

func SetListPageMaxNum(ctx context.Context, v int64) context.Context {
	return context.WithValue(ctx, ListPageMaxNumKey, v)
}

func SetAdministratorAuthData(ctx context.Context, v *AdministratorAuthData) context.Context {
	return context.WithValue(ctx, AdministratorAuthDataKey, v)
}

func SetVisitorAuthData(ctx context.Context, v *VisitorAuthData) context.Context {
	return context.WithValue(ctx, VisitorAuthDataKey, v)
}

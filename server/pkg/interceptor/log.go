package interceptor

import (
	"context"
	"fmt"
	"time"
	"yuumi/pkg/crypto"
	"yuumi/pkg/keys"
	"yuumi/pkg/logger"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GetLogUnaryInterceptor(log *logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = withUUID(ctx)
		ctx = keys.SetRequestFullMethod(ctx, info.FullMethod)
		ctx = keys.SetRequestStartTime(ctx, time.Now())

		defer func() {
			kv := getRequestLogData(ctx, req, info)
			log.Get("Info").Sugar().Infow(info.FullMethod, kv...)
		}()

		return handler(ctx, req)
	}
}

func withUUID(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)

	var requestId string
	if v := md.Get("request-id"); len(v) > 0 && v[0] != "" {
		requestId = v[0]
	} else {
		requestId = uuid.New().String()
	}

	return keys.SetRequestID(ctx, requestId)
}

func getRequestLogData(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) []interface{} {
	requestId := keys.GetRequestIDFormContext(ctx)
	md, _ := metadata.FromIncomingContext(ctx)
	kv := []interface{}{
		"requestID", requestId,
		"duration(ms)", time.Since(keys.GetRequestStartTimeFromContext(ctx)).Milliseconds(),
		"request", crypto.AesEncrypt(fmt.Sprintf("%#v", req), requestId),
	}

	// http request method
	if v := md.Get("grpcgateway-method"); v != nil {
		kv = append(kv, "httpMethod", v[0])
	}
	// http request path
	if v := md.Get("grpcgateway-path"); v != nil {
		kv = append(kv, "httpPath", v[0])
	}

	// http user agent
	if v := md.Get("grpcgateway-user-agent"); v != nil {
		kv = append(kv, "userAgent", v[0])
	} else if v := md.Get("user-agent"); v != nil {
		kv = append(kv, "userAgent", v)
	}

	return kv
}

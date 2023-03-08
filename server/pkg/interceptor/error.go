package interceptor

import (
	"context"
	"yuumi/pkg/keys"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func GetErrorLogUnaryInterceptor(log *logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		res, err := handler(ctx, req)

		if err != nil {
			kv := []interface{}{
				"requestID", keys.GetRequestIDFormContext(ctx),
				"err", err,
			}

			defer func(err error) {
				log.Get("Error").Sugar().Errorw(info.FullMethod, kv...)
			}(err)

			// 将details信息载入log中，同时在返回值中去除
			if s, ok := status.FromError(err); ok {
				kv = append(kv, "details", s.Details())
				err = status.Error(s.Code(), s.Message())
			}
		}

		return res, err
	}
}

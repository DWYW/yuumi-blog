package interceptor

import (
	"context"
	"regexp"
	"yuumi/pkg/keys"

	"google.golang.org/grpc"
)

func GetListUnaryInterceptor(maxNum int64) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 分页列表注入最大获取量
		if ok, _ := regexp.MatchString("GetList$", info.FullMethod); ok {
			if maxNum <= 0 || maxNum >= 20 {
				maxNum = int64(20)
			}
			ctx = keys.SetListPageMaxNum(ctx, maxNum)
		}

		return handler(ctx, req)
	}
}

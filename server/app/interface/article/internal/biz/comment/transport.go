package comment

import (
	"context"
	v1 "yuumi/api/interface/article/v1"
	"yuumi/pkg/errorcode"
)

func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

func DecodeCreateRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.CreateCommentRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}

func DecodeGetListRequest(_ context.Context, r interface{}) (interface{}, error) {
	request, ok := r.(*v1.GetCommentListRequest)
	if !ok {
		return nil, errorcode.New(errorcode.ArgumentInvalid)
	}
	return request, nil
}

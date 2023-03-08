package errorcode

import (
	"google.golang.org/grpc/status"
)

func New(code Code) error {
	return NewWithDetail(code, code.ErrMesssage)
}

func NewWithDetail(code Code, detail interface{}) error {
	var message string

	if v, ok := detail.(string); ok {
		message = v
	} else if v, ok := detail.(error); ok {
		message = v.Error()
	}

	stat, _ := status.New(code.StatusCode, code.ErrMesssage).WithDetails(&ErrorDetail{
		ErrCode: code.ErrCode,
		ErrMsg:  message,
	})

	return stat.Err()
}

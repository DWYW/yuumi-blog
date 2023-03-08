package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"yuumi/pkg/keys"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type GrpcClient struct {
	Host string
	Port uint64
}

func (c GrpcClient) NewClient() (*GrpcRequestClient, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", c.Host, c.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		bytes, _ := json.Marshal(struct {
			ErrCode int64  `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
		}{ErrCode: 500, ErrMsg: "service connenct err"})

		return nil, errors.New(string(bytes))
	}

	return &GrpcRequestClient{Conn: conn}, nil
}

type GrpcRequestClient struct {
	Conn     *grpc.ClientConn
	handlers []func(ctx context.Context, cc grpc.ClientConnInterface) error
}

func (c *GrpcRequestClient) Handle(handler func(ctx context.Context, cc grpc.ClientConnInterface) error) *GrpcRequestClient {
	c.handlers = append(c.handlers, handler)
	return c
}

func (c *GrpcRequestClient) background(ctx context.Context) context.Context {
	return metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"request-id": keys.GetRequestIDFormContext(ctx),
	}))
}

func (c *GrpcRequestClient) Exec(ctx context.Context) error {
	defer func() {
		c.Conn.Close()
	}()

	var err error
	bg := c.background(ctx)
	for _, handler := range c.handlers {
		if err = handler(bg, c.Conn); err != nil {
			break
		}
	}

	return err
}

func (c *GrpcRequestClient) Request(ctx context.Context, handler func(ctx context.Context, cc grpc.ClientConnInterface) (interface{}, error)) (interface{}, error) {
	return handler(c.background(ctx), c.Conn)
}

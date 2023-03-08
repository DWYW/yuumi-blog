package task

import (
	"context"
	"yuumi/pkg/client"

	"log"

	v1 "yuumi/api/service/administrator/v1"

	"google.golang.org/grpc"
)

type PermissionData struct {
	Name      string
	RpcMethod string
}

func (t Task) CreatePermissions(permissions []PermissionData) []int64 {
	c, err := client.GrpcClient{Host: t.Host, Port: t.Port}.NewClient()
	if err != nil {
		log.Println(err)
	}

	var data []int64

	for i := 0; i < len(permissions); i++ {
		v := permissions[i]

		c.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
			client := v1.NewPermissionServiceClient(cc)

			info, err := client.GetInfo(ctx, &v1.GetPermissionInfoRequest{RpcMethod: v.RpcMethod})
			if err == nil {
				data = append(data, info.Permission.Id)
			} else {
				created, err := client.Create(ctx, &v1.CreatePermissionRequest{Name: v.Name, RpcMethod: v.RpcMethod})
				if err == nil {
					data = append(data, created.Permission.Id)
				}
			}

			return nil
		})
	}

	c.Exec(context.Background())
	return data
}

package task

import (
	"context"
	"yuumi/pkg/client"

	"log"

	v1 "yuumi/api/service/administrator/v1"

	"google.golang.org/grpc"
)

type RoleData struct {
	Name        string
	Type        string
	Permissions []PermissionData
	Navs        []NavData
	Children    []RoleData
}

func (t Task) CreateRoles(roles []RoleData, parentId int64) []int64 {
	c, err := client.GrpcClient{Host: t.Host, Port: t.Port}.NewClient()
	if err != nil {
		log.Println(err)
	}

	var data []int64

	for i := 0; i < len(roles); i++ {
		v := roles[i]
		c.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
			role := &v1.RoleData{}
			client := v1.NewRoleServiceClient(cc)
			created, err := client.Create(ctx, &v1.CreateRoleRequest{Name: v.Name, Type: v.Type, ParentId: parentId})
			if err == nil {
				role = created.Role
			} else {
				info, err := client.GetInfo(ctx, &v1.GetRoleInfoRequest{Type: v.Type})
				if err == nil {
					role = info.Role
				}
			}

			if role != nil {
				// 只返回顶层Role
				data = append(data, role.Id)

				// 创建权限
				ids := t.CreatePermissions(v.Permissions)
				client.AppendPermissionsWithRoleID(ctx, &v1.AppendPermissionsWithRoleIDRequest{RoleId: role.Id, PermissionIds: ids})

				if len(v.Children) > 0 {
					t.CreateRoles(v.Children, role.Id)
				}

				// 创建菜单
				ids = t.CreateNavs(v.Navs, 0)
				t.BindWithRoleId(ids, role.Id)
			}

			return nil
		})
	}

	c.Exec(context.Background())
	return data
}

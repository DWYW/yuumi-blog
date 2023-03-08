package task

import (
	"context"
	"yuumi/pkg/client"

	"log"

	v1 "yuumi/api/service/administrator/v1"

	"google.golang.org/grpc"
)

type AdministratorData struct {
	Name     string
	Password string
	Roles    []RoleData
}

func (t Task) CreateAdministrators(administrators []AdministratorData) {
	c, err := client.GrpcClient{Host: t.Host, Port: t.Port}.NewClient()
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(administrators); i++ {
		v := administrators[i]

		c.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
			client := v1.NewAdministratorServiceClient(cc)

			var administrator *v1.AdministratorData
			created, err := client.Create(ctx, &v1.CreateAdministratorRequest{Name: v.Name, Password: v.Password})
			if err != nil {
				info, err := client.GetAdministratorWithNameAndPassword(ctx, &v1.GetAdministratorWithNameAndPasswordRequest{Name: v.Name, Password: v.Password})
				if err == nil {
					administrator = info.Administrator
				}
			} else {
				administrator = created.Administrator
			}

			// 创建关联角色
			roleIds := t.CreateRoles(v.Roles, 0)

			// 绑定角色
			client.AppendRolesWithAdministratorID(ctx, &v1.AppendRolesWithAdministratorIDRequest{
				AdministratorId: administrator.Id,
				RoleIds:         roleIds,
			})

			return err
		})
	}

	c.Exec(context.Background())
}

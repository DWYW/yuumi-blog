package task

import (
	"context"
	"log"
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/pkg/client"

	"google.golang.org/grpc"
)

type NavData struct {
	Icon        string
	Name        string
	LinkUrl     string
	ActivedRule string
	Weight      int64
	Children    []NavData
}

func (t Task) CreateNavs(navs []NavData, parentId int64) []int64 {
	c, err := client.GrpcClient{Host: t.Host, Port: t.Port}.NewClient()
	if err != nil {
		log.Println(err)
	}

	var data []int64

	for i := 0; i < len(navs); i++ {
		v := navs[i]
		c.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
			nav := &v1.NavMenuData{}
			client := v1.NewNavMenuServiceClient(cc)

			info, err := client.GetInfo(ctx, &v1.GetNavMenuInfoRequest{Name: v.Name})
			if err == nil {
				data = append(data, info.Menu.Id)
				nav = info.Menu
			} else {
				created, err := client.Create(ctx, &v1.CreateNavMenuRequest{
					Name:        v.Name,
					Icon:        &v.Icon,
					LinkUrl:     &v.LinkUrl,
					ActivedRule: &v.ActivedRule,
					Weight:      &v.Weight,
					ParentId:    &parentId,
				})

				if err == nil {
					data = append(data, created.Menu.Id)
					nav = created.Menu
				}
			}

			if len(v.Children) > 0 {
				children := t.CreateNavs(v.Children, nav.Id)
				data = append(data, children...)
			}

			return nil
		})
	}

	c.Exec(context.Background())
	return data
}

func (t Task) BindWithRoleId(navIds []int64, roleId int64) {
	c, err := client.GrpcClient{Host: t.Host, Port: t.Port}.NewClient()
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(navIds); i++ {
		v := navIds[i]
		c.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
			client := v1.NewNavMenuServiceClient(cc)

			client.BindWithRoleIDs(ctx, &v1.NavMenuBindWithRoleIDsRequest{
				Id:      v,
				RoleIds: []int64{roleId},
			})

			return nil
		})
	}

	c.Exec(context.Background())
}

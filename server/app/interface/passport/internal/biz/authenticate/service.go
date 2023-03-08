package authenticate

import (
	"context"
	"strings"
	v1 "yuumi/api/interface/passport/v1"
	administratorV1 "yuumi/api/service/administrator/v1"
	"yuumi/app/interface/passport/internal/conf"
	"yuumi/app/interface/passport/internal/pkg/jwt"
	"yuumi/pkg/client"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/keys"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

type ServiceInterface interface {
	AdministratorWithJWT(ctx context.Context, in *v1.AdministratorWithJWTAuthenticateRequest) (*v1.AdministratorWithJWTAuthenticateReply, error)
	VisitorWithJWT(ctx context.Context, in *v1.VisitorWithJWTAuthenticateRequest) (*v1.VisitorWithJWTAuthenticateReply, error)
}

type Service struct {
	Log                 *logger.Logger
	AdministratorClient *client.GrpcClient
}

func (s Service) AdministratorWithJWT(ctx context.Context, in *v1.AdministratorWithJWTAuthenticateRequest) (*v1.AdministratorWithJWTAuthenticateReply, error) {
	var token = strings.TrimPrefix(keys.GetAuthorizationFromContext(ctx), "Bearer ")
	jwtConf := conf.GetJwtConfig().Administrator
	administratorJWT := jwt.AdministratorJWT{Iv: jwtConf.Iv, Secret: jwtConf.Key}
	claims, err := administratorJWT.Parse(token)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.AuthorizationInvalid, err)
	}

	if !claims.VerifyAudience(jwt.AdministratorAudience, false) {
		return nil, errorcode.New(errorcode.AuthorizationInvalid)
	}

	// 鉴定方法是否在权限内
	if in.Method != "" {
		client, err := s.AdministratorClient.NewClient()
		if err != nil {
			return nil, err
		}

		var (
			permission *administratorV1.GetPermissionInfoReply
			target     *administratorV1.GetRoleInfoWithPermissionIDReply
			sources    *administratorV1.GetRolesWithAdministratorIDReply
		)

		client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
			client := administratorV1.NewPermissionServiceClient(cc)
			permission, err = client.GetInfo(ctx, &administratorV1.GetPermissionInfoRequest{RpcMethod: in.Method})
			return err
		}).Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
			client := administratorV1.NewRoleServiceClient(cc)
			target, err = client.GetInfoWithPermissionID(ctx, &administratorV1.GetRoleInfoWithPermissionIDRequest{PermissionId: permission.Permission.Id})
			return err
		}).Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
			client := administratorV1.NewRoleServiceClient(cc)
			sources, err = client.GetRolesWithAdministratorID(ctx, &administratorV1.GetRolesWithAdministratorIDRequest{AdministratorId: claims.AdministratorCustomInfo.ID})
			return err
		})

		if err = client.Exec(ctx); err != nil {
			return nil, errorcode.New(errorcode.PermissionDenied)
		}

		var ok bool
		for _, role := range sources.Roles {
			if ok = role.Id == target.Role.Id; ok {
				break
			}
		}
		if !ok {
			return nil, errorcode.New(errorcode.PermissionDenied)
		}
	}

	return &v1.AdministratorWithJWTAuthenticateReply{
		Id:   claims.AdministratorCustomInfo.ID,
		Name: claims.AdministratorCustomInfo.Name,
	}, nil
}

func (s Service) VisitorWithJWT(ctx context.Context, in *v1.VisitorWithJWTAuthenticateRequest) (*v1.VisitorWithJWTAuthenticateReply, error) {
	var token = strings.TrimPrefix(keys.GetAuthorizationFromContext(ctx), "Bearer ")
	jwtConf := conf.GetJwtConfig().Visitor
	visitorJwt := jwt.VisitorJWT{Iv: jwtConf.Iv, Secret: jwtConf.Key}
	claims, err := visitorJwt.Parse(token)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.AuthorizationInvalid, err)
	}

	if !claims.VerifyAudience(jwt.VisitorAudience, false) {
		return nil, errorcode.New(errorcode.AuthorizationInvalid)
	}

	return &v1.VisitorWithJWTAuthenticateReply{
		Id:   claims.VisitorTokenInfo.ID,
		Name: claims.VisitorTokenInfo.Name,
	}, nil
}

package authorize

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "yuumi/api/interface/passport/v1"
	administratorV1 "yuumi/api/service/administrator/v1"
	articleV1 "yuumi/api/service/article/v1"

	"yuumi/app/interface/passport/internal/conf"
	"yuumi/app/interface/passport/internal/pkg/jwt"
	"yuumi/pkg/client"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

type ServiceInterface interface {
	AdministratorWithPassword(ctx context.Context, in *v1.AdministratorWithPasswordAuthorizeRequest) (*v1.AdministratorWithPasswordAuthorizeReply, error)
	VisitorWithGithubSession(ctx context.Context, in *v1.VisitorWithGithubSessionAuthorizeRequest) (*v1.VisitorWithGithubSessionAuthorizeReply, error)
}

type Service struct {
	Log                 *logger.Logger
	AdministratorClient *client.GrpcClient
	ArticleClient       *client.GrpcClient
}

func (s Service) AdministratorWithPassword(ctx context.Context, in *v1.AdministratorWithPasswordAuthorizeRequest) (*v1.AdministratorWithPasswordAuthorizeReply, error) {
	client, err := s.AdministratorClient.NewClient()
	if err != nil {
		return nil, err
	}

	// 通过用户名和密码获取信息
	var administratorReply *administratorV1.GetAdministratorWithNameAndPasswordReply
	client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := administratorV1.NewAdministratorServiceClient(cc)
		administratorReply, err = client.GetAdministratorWithNameAndPassword(ctx, &administratorV1.GetAdministratorWithNameAndPasswordRequest{
			Name:     in.Name,
			Password: in.Password,
		})
		return err
	})

	if err = client.Exec(ctx); err != nil {
		return nil, err
	}

	// 创建Administrator JWT
	jwtConf := conf.GetJwtConfig().Administrator
	administratorJWT := jwt.AdministratorJWT{Iv: jwtConf.Iv, Secret: jwtConf.Key}
	token, err := administratorJWT.Signed(&jwt.AdministratorCustomInfo{
		ID:   administratorReply.Administrator.Id,
		Name: administratorReply.Administrator.Name,
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.TokenCreateFailed, err)
	}

	return &v1.AdministratorWithPasswordAuthorizeReply{
		Token:         token,
		Administrator: administratorReply.Administrator,
	}, nil
}

type UserResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	HtmlUrl   string `json:"html_url"`
	Blog      string `json:"blog"`
}

func (s Service) VisitorWithGithubSession(ctx context.Context, in *v1.VisitorWithGithubSessionAuthorizeRequest) (*v1.VisitorWithGithubSessionAuthorizeReply, error) {
	// 解析session
	sessionConf := conf.GetJwtConfig().VisitorSession
	visitorSession := jwt.VisitorSession{Iv: sessionConf.Iv, Secret: sessionConf.Key}
	claims, err := visitorSession.Parse(in.Session)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.SessionInvalid, err)
	}

	if !claims.VerifyAudience(jwt.VisitorAudience, false) {
		return nil, errorcode.New(errorcode.SessionInvalid)
	}

	info := claims.VisitorSessionInfo

	// 通过session token 获取GitHub 用户数据
	ins := client.HttpClient{}.NewClient()
	res, err := ins.Request(client.RequestOption{
		Method: "GET",
		Path:   "https://api.github.com/user",
		Headers: &map[string]string{
			"Authorization": fmt.Sprintf("%s %s", info.TokenType, info.Token),
		},
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.Unknown, err)
	}

	var githubUser UserResponse
	json.Unmarshal(res, &githubUser)

	// 通过GitHubID 查找
	rpc, err := s.ArticleClient.NewClient()
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.Unknown, err)
	}

	temp, err := rpc.Request(ctx, func(ctx context.Context, cc grpc.ClientConnInterface) (interface{}, error) {
		c := articleV1.NewVisitorServiceClient(cc)
		return c.GetVisitorWithGithubID(ctx, &articleV1.GetVisitorWithGithubIDRequest{GithubId: githubUser.ID})
	})
	if err != nil {
		return nil, err
	}

	var visitor *articleV1.VisitorData
	if v, _ := temp.(*articleV1.GetVisitorWithGithubIDReply); v != nil {
		visitor = v.Visitor
	}

	githubUserUpdate := articleV1.GithubUser{
		GithubId:  githubUser.ID,
		Name:      githubUser.Name,
		AvatarUrl: githubUser.AvatarUrl,
		HtmlUrl:   githubUser.HtmlUrl,
		Blog:      githubUser.Blog,
	}
	// 根据结果创建或更新
	if visitor == nil {
		reply, err := rpc.Request(ctx, func(ctx context.Context, cc grpc.ClientConnInterface) (interface{}, error) {
			c := articleV1.NewVisitorServiceClient(cc)
			return c.CreateWithGithubUser(ctx, &articleV1.CreateWithGithubUserRequest{GithubUser: &githubUserUpdate})
		})

		if err != nil {
			return nil, err
		}

		visitor = reply.(*articleV1.CreateWithGithubUserReply).Visitor
	} else {
		_, err := rpc.Request(ctx, func(ctx context.Context, cc grpc.ClientConnInterface) (interface{}, error) {
			c := articleV1.NewVisitorServiceClient(cc)
			return c.UpdateWithGithubUser(ctx, &articleV1.UpdateWithGithubUserRequest{Id: visitor.Id, GithubUser: &githubUserUpdate})
		})

		if err != nil {
			return nil, err
		}
	}

	// 创建Token
	jwtConf := conf.GetJwtConfig().Visitor
	visitorJwt := jwt.VisitorJWT{Iv: jwtConf.Iv, Secret: jwtConf.Key}
	token, err := visitorJwt.Signed(&jwt.VisitorTokenInfo{
		ID:   visitor.Id,
		Name: visitor.Name,
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.TokenCreateFailed, err)
	}

	return &v1.VisitorWithGithubSessionAuthorizeReply{Token: token, Visitor: visitor}, nil
}

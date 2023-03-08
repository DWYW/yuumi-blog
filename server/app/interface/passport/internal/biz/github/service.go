package github

import (
	"context"
	"encoding/json"
	v1 "yuumi/api/interface/passport/v1"
	"yuumi/app/interface/passport/internal/conf"
	"yuumi/app/interface/passport/internal/pkg/jwt"
	"yuumi/pkg/errorcode"

	"yuumi/pkg/client"
	"yuumi/pkg/logger"
)

type ServiceInterface interface {
	GetSessionWithGithubCode(ctx context.Context, in *v1.GetSessionWithGithubCodeRequest) (*v1.GetSessionWithGithubCodeReply, error)
}

type Service struct {
	Log                *logger.Logger
	GithubClient       *client.HttpClient
	GithubClientID     string
	GithubClientSecret string
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	Error       string `json:"error"`
}

func (s Service) GetSessionWithGithubCode(ctx context.Context, in *v1.GetSessionWithGithubCodeRequest) (*v1.GetSessionWithGithubCodeReply, error) {
	// code换token
	res, err := s.GithubClient.NewClient().Request(client.RequestOption{
		Duration: 10,
		Method:   "POST",
		Path:     "https://github.com/login/oauth/access_token",
		Data: &map[string]string{
			"client_id":     s.GithubClientID,
			"client_secret": s.GithubClientSecret,
			"code":          in.Code,
			"redirect_uri":  in.Redirect,
		},
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.Unknown, err)
	}

	var reply AccessTokenResponse
	json.Unmarshal(res, &reply)

	if reply.Error != "" {
		return nil, errorcode.NewWithDetail(errorcode.Unknown, reply.Error)
	}

	// 创建session
	jwtConf := conf.GetJwtConfig().VisitorSession
	visitorSession := jwt.VisitorSession{Iv: jwtConf.Iv, Secret: jwtConf.Key}
	session, err := visitorSession.Signed(&jwt.VisitorSessionInfo{
		Token:     reply.AccessToken,
		TokenType: reply.TokenType,
		Scope:     reply.Scope,
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.SessionCreateFailed, err)
	}

	return &v1.GetSessionWithGithubCodeReply{Session: session}, nil
}

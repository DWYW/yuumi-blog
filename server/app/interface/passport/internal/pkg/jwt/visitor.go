package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type VisitorSessionInfo struct {
	Token     string
	TokenType string
	Scope     string
}

type VisitorSessionClaims struct {
	*jwt.StandardClaims
	*VisitorSessionInfo
}

type VisitorSession struct {
	Iv     uint64
	Secret string
}

func (t *VisitorSession) Signed(info *VisitorSessionInfo) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = &VisitorSessionClaims{
		&jwt.StandardClaims{
			Audience:  VisitorAudience,                                          // 受众群体
			ExpiresAt: time.Now().Add(time.Second * time.Duration(t.Iv)).Unix(), // 到期时间
			Id:        "",                                                       // 编号
			IssuedAt:  time.Now().Unix(),                                        // 签发时间
			Issuer:    "",                                                       // 签发人
			NotBefore: time.Now().Unix(),                                        // 生效时间
			Subject:   "",                                                       // 主题
		},
		info,
	}

	return token.SignedString([]byte(t.Secret))
}

func (t *VisitorSession) Parse(token string) (*VisitorSessionClaims, error) {
	result, err := jwt.ParseWithClaims(token, &VisitorSessionClaims{}, func(v *jwt.Token) (interface{}, error) {
		return []byte(t.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, _ := result.Claims.(*VisitorSessionClaims)

	return claims, nil
}

type VisitorTokenInfo struct {
	ID   int64
	Name string
}

type VisitorTokenClaims struct {
	*jwt.StandardClaims
	*VisitorTokenInfo
}

type VisitorJWT struct {
	Iv     uint64
	Secret string
}

func (t *VisitorJWT) Signed(info *VisitorTokenInfo) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = &VisitorTokenClaims{
		&jwt.StandardClaims{
			Audience:  VisitorAudience,                                          // 受众群体
			ExpiresAt: time.Now().Add(time.Second * time.Duration(t.Iv)).Unix(), // 到期时间
			Id:        "",                                                       // 编号
			IssuedAt:  time.Now().Unix(),                                        // 签发时间
			Issuer:    "",                                                       // 签发人
			NotBefore: time.Now().Unix(),                                        // 生效时间
			Subject:   "",                                                       // 主题
		},
		info,
	}

	return token.SignedString([]byte(t.Secret))
}

func (t *VisitorJWT) Parse(token string) (*VisitorTokenClaims, error) {
	result, err := jwt.ParseWithClaims(token, &VisitorTokenClaims{}, func(v *jwt.Token) (interface{}, error) {
		return []byte(t.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, _ := result.Claims.(*VisitorTokenClaims)

	return claims, nil
}

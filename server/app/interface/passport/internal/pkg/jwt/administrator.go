package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AdministratorCustomInfo struct {
	ID   int64
	Name string
}

type AdministratorTokenClaims struct {
	*jwt.StandardClaims
	*AdministratorCustomInfo
}

type AdministratorJWT struct {
	Iv     uint64
	Secret string
}

func (t *AdministratorJWT) Signed(info *AdministratorCustomInfo) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = &AdministratorTokenClaims{
		&jwt.StandardClaims{
			Audience:  AdministratorAudience,                                    // 受众群体
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

func (t *AdministratorJWT) Parse(token string) (*AdministratorTokenClaims, error) {
	result, err := jwt.ParseWithClaims(token, &AdministratorTokenClaims{}, func(v *jwt.Token) (interface{}, error) {
		return []byte(t.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, _ := result.Claims.(*AdministratorTokenClaims)

	return claims, nil
}

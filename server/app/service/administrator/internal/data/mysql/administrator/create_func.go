package administrator

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/pkg/crypto"
)

func (model *Model) PasswordValidator(password string, target string, salt string) bool {
	return crypto.HmacSha256(salt, password) == target
}

func (model *Model) CreateOne(ctx context.Context, v *schema.Administrator) (*schema.Administrator, error) {
	// 密码加盐
	salt, err := crypto.GetSlat(16)
	if err != nil {
		return nil, err
	}

	v.Salt = salt
	v.Password = crypto.HmacSha256(salt, v.Password)

	result := model.GetDB().WithContext(ctx).Create(v)
	if err := result.Error; err != nil {
		return nil, err
	}

	return model.FindOne(ctx, &FindCondition{ID: int64(v.ID)})
}

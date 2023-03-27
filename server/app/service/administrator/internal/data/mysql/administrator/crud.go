// 此文件存在基本的CRUD通用逻辑，特殊逻辑请存放到其他文件下
package administrator

import (
	"context"
	"fmt"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/pkg/crypto"

	"gorm.io/gorm"
)

/*
 *  创建记录
 */
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

	return model.FindOne(ctx, &FindCondition{
		Where: &FindConditionWhere{ID: int64(v.ID)},
	})
}

/*
 *  查找单条记录
 */

func (model *Model) FindOne(ctx context.Context, condition *FindCondition) (*schema.Administrator, error) {
	dest := schema.Administrator{}
	result := model.GetDB().WithContext(ctx).Scopes(condition.Build()...).First(&dest)
	return &dest, result.Error
}

func (model *Model) FindByID(ctx context.Context, id int64) (*schema.Administrator, error) {
	return model.FindOne(ctx, &FindCondition{
		Where: &FindConditionWhere{ID: id},
	})
}

/*
 *  查找多条记录
 */

func (model *Model) Find(ctx context.Context, condition *FindCondition, scopes ...func(*gorm.DB) *gorm.DB) ([]*schema.Administrator, error) {
	var dest []*schema.Administrator
	tx := model.GetDB().WithContext(ctx)
	result := tx.Scopes(scopes...).Scopes(condition.Build()...).Find(&dest)
	return dest, result.Error
}

func (model *Model) FindList(ctx context.Context, condition *FindListCondition) ([]*schema.Administrator, error) {
	skipNum := (condition.Page - 1) * condition.PageSize
	return model.Find(ctx, condition.FindCondition, func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(skipNum)).Limit(int(condition.PageSize))
	})
}

/*
 *  统计数量
 */

func (model *Model) Count(ctx context.Context, where *FindConditionWhere, exludeDeleted bool) (int64, error) {
	scopes := where.Build()
	if exludeDeleted {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` IS NULL", model.Fields.DeletedAt))
		})
	}

	var count int64
	tx := model.GetDB().Table(model.TableName).WithContext(ctx)
	result := tx.Scopes(scopes...).Count(&count)
	return count, result.Error
}

/*
 *  根据ID更新记录
 */

func (model *Model) UpdateWithID(ctx context.Context, id int64, v interface{}) (*schema.Administrator, error) {
	administrator := schema.Administrator{Model: gorm.Model{ID: uint(id)}}

	result := model.GetDB().WithContext(ctx).Model(&administrator).Omit(model.Fields.Salt).Updates(v)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.FindOne(ctx, &FindCondition{
		Where: &FindConditionWhere{ID: id},
	})
}

/*
 *  删除单条记录
 */

func (model *Model) DeleteOne(ctx context.Context, condition *DeleteCondition) (*schema.Administrator, error) {
	dest := schema.Administrator{}
	result := model.GetDB().WithContext(ctx).Scopes(condition.Build()...).Delete(&dest)
	return &dest, result.Error
}

func (model *Model) DeleteWithID(ctx context.Context, id int64) (*schema.Administrator, error) {
	condition := DeleteCondition{ID: id}
	return model.DeleteOne(ctx, &condition)
}

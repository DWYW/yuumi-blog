// 此文件存在基本的CRUD通用逻辑的条件信息，特殊逻辑条件信息请存放到其他文件下
package comment

import (
	"fmt"
	"reflect"
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func forEach(c interface{}, callback func(key string, value interface{})) {
	v := reflect.ValueOf(c).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		value := v.Field(i).Interface()
		key := t.Field(i).Name
		callback(key, value)
	}
}

/*
 *  FindCondition
 */

type FindCondition struct {
	Where   *FindConditionWhere
	Sort    *FindConditionSort
	Preload *FindConditionPreload
}

func (condition *FindCondition) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)

	// 筛选条件
	if condition.Where != nil {
		scopes = append(scopes, condition.Where.Build()...)
	}

	// 预加载条件
	if condition.Preload != nil {
		scopes = append(scopes, condition.Preload.Build()...)
	}

	// 排序条件
	if condition.Sort != nil {
		scopes = append(scopes, condition.Sort.Build()...)
	}

	return scopes
}

type FindConditionWhere struct {
	ID        int64
	ArticleID int64
	CreatorID int64
}

func (condition *FindConditionWhere) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	fields := schema.Comment{}.GetFields()

	if condition.ID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ID), condition.ID)
		})
	}

	if condition.ArticleID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ArticleID), condition.ArticleID)
		})
	}

	if condition.CreatorID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.CreatorID), condition.CreatorID)
		})
	}

	return scopes
}

type FindConditionPreload struct {
	Article bool
	Creator bool
}

func (condition *FindConditionPreload) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	relations := schema.Comment{}.GetRelations()
	rv := reflect.ValueOf(relations).Elem()
	forEach(condition, func(key string, value interface{}) {
		relation := rv.FieldByName(key).String()
		if ok := value.(bool); ok && relation != "" {
			scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
				return db.Preload(relation)
			})
		}
	})

	return scopes
}

type FindConditionSort struct {
	ID int64
}

func (condition *FindConditionSort) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	fields := schema.Comment{}.GetFields()
	rv := reflect.ValueOf(fields).Elem()
	forEach(condition, func(key string, value interface{}) {
		field := rv.FieldByName(key).String()
		if v := value.(int64); v != 0 && field != "" {
			sortType := "desc"
			if v > 0 {
				sortType = "asc"
			}

			scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
				return db.Order(fmt.Sprintf("%s %s", field, sortType))
			})
		}
	})

	return scopes
}

/*
 *  FindListCondition
 */

type FindListCondition struct {
	Page          int64
	PageSize      int64
	FindCondition *FindCondition
}

/*
 *  DeleteCondition
 */
type DeleteCondition struct {
	ID int64
}

func (condition *DeleteCondition) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	fields := schema.Comment{}.GetFields()

	if condition.ID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ID), condition.ID)
		})
	}

	return scopes
}

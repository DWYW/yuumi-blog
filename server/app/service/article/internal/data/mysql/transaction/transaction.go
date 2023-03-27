package transaction

import (
	"yuumi/app/service/article/internal/data/mysql/article"
	"yuumi/app/service/article/internal/data/mysql/comment"
	"yuumi/app/service/article/internal/data/mysql/githubuser"
	"yuumi/app/service/article/internal/data/mysql/visitor"

	"gorm.io/gorm"
)

type Transaction struct {
	DB            *gorm.DB
	GetArticle    func(databases ...*gorm.DB) *article.Model
	GetVisitor    func(databases ...*gorm.DB) *visitor.Model
	GetGithubUser func(databases ...*gorm.DB) *githubuser.Model
	GetComment    func(databases ...*gorm.DB) *comment.Model
}

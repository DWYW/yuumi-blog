package mysql

import (
	"yuumi/app/service/article/internal/data/mysql/article"
	"yuumi/app/service/article/internal/data/mysql/comment"
	"yuumi/app/service/article/internal/data/mysql/githubuser"
	"yuumi/app/service/article/internal/data/mysql/transaction"
	"yuumi/app/service/article/internal/data/mysql/visitor"

	"gorm.io/gorm"
)

func GetClient() *gorm.DB {
	return client
}

func getDefaultDB(databases ...*gorm.DB) *gorm.DB {
	var db *gorm.DB
	for _, item := range databases {
		if item != nil {
			db = item
			break
		}
	}

	if db == nil {
		db = client
	}

	return db
}

func GetArticle(databases ...*gorm.DB) *article.Model {
	return article.New(getDefaultDB(databases...))
}

func GetVisitor(databases ...*gorm.DB) *visitor.Model {
	return visitor.New(getDefaultDB(databases...))
}

func GetGithubUser(databases ...*gorm.DB) *githubuser.Model {
	return githubuser.New(getDefaultDB(databases...))
}

func GetComment(databases ...*gorm.DB) *comment.Model {
	return comment.New(getDefaultDB(databases...))
}

func GetTransaction(databases ...*gorm.DB) *transaction.Transaction {
	db := getDefaultDB(databases...)
	return &transaction.Transaction{
		DB:            db,
		GetArticle:    GetArticle,
		GetVisitor:    GetVisitor,
		GetGithubUser: GetGithubUser,
		GetComment:    GetComment,
	}
}

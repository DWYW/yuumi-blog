package schema

import "gorm.io/gorm"

type CommentFields struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Content   string
	CreatorID string
	ArticleID string
}

type CommentRelations struct {
	Creator string
	Article string
}

type Comment struct {
	gorm.Model
	Content   string   `gorm:"default:'';comment:简述"`
	CreatorID int64    `gorm:"comment:创建人ID"`
	Creator   *Visitor `gorm:"foreignKey:CreatorID;references:ID"`
	ArticleID int64    `gorm:"comment:文章ID"`
	Article   *Article
}

func (Comment) TableName() string {
	return "comments"
}

func (Comment) GetFields() *CommentFields {
	return &CommentFields{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		Content:   "content",
		CreatorID: "creator_id",
		ArticleID: "article_id",
	}
}

func (Comment) GetRelations() *CommentRelations {
	return &CommentRelations{
		Creator: "Creator",
		Article: "Article",
	}
}

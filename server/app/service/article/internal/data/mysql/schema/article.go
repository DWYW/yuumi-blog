package schema

import "gorm.io/gorm"

type ArticleFields struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	Title       string
	Description string
	Keyword     string
	Content     string
	CoverUrl    string
}

type Article struct {
	gorm.Model
	Title       string
	Description string `gorm:"size:255;default:'';comment:简述"`
	Keyword     string `gorm:"size:255;default:'';comment:文章关键词"`
	Content     string
	CoverUrl    string `gorm:"size:255;default:'';comment:封面图链接地址"`
}

func (Article) GetFeilds() *ArticleFields {
	return &ArticleFields{
		ID:          "id",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
		Title:       "title",
		Description: "description",
		Keyword:     "keyword",
		Content:     "content",
		CoverUrl:    "cover_url",
	}
}

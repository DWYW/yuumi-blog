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
	Status      string
}

type Article struct {
	gorm.Model
	Title       string
	Description string `gorm:"size:255;default:'';comment:简述"`
	Keyword     string `gorm:"size:255;default:'';comment:文章关键词"`
	Content     string
	CoverUrl    string `gorm:"size:255;default:'';comment:封面图链接地址"`
	Status      int64  `gorm:"default:0;comment:0:未发布,1:已发布"`
}

func (Article) TableName() string {
	return "articles"
}

func (Article) GetFields() *ArticleFields {
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
		Status:      "status",
	}
}

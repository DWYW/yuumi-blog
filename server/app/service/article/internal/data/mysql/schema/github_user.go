package schema

import "gorm.io/gorm"

type GithubUserFields struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	VisitorID string
	GithubID  string
	Name      string
	AvatarUrl string
	Blog      string
	HtmlUrl   string
}

type GithubUser struct {
	gorm.Model
	VisitorID uint64
	GithubID  uint64
	Name      string
	AvatarUrl string `gorm:"size:255;default:'';comment:头像链接"`
	Blog      string `gorm:"size:255;default:'';comment:博客地址"`
	HtmlUrl   string `gorm:"size:255;default:'';comment:Github地址"`
}

func (GithubUser) TableName() string {
	return "github_users"
}

func (GithubUser) GetFields() *GithubUserFields {
	return &GithubUserFields{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		VisitorID: "visitor_id",
		GithubID:  "github_id",
		Name:      "name",
		AvatarUrl: "avatar_url",
		Blog:      "blog",
		HtmlUrl:   "html_url",
	}
}

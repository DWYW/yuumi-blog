package schema

import "gorm.io/gorm"

type VisitorFields struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
	AvatarUrl string
}

type Visitor struct {
	gorm.Model
	Name       string
	AvatarUrl  string `gorm:"size:255;default:'';comment:头像链接"`
	GithubUser GithubUser
}

func (Visitor) GetFeilds() *VisitorFields {
	return &VisitorFields{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		Name:      "name",
		AvatarUrl: "avatar_url",
	}
}

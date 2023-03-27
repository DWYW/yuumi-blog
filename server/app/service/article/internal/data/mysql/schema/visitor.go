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

type VisitorRelations struct {
	GithubUser string
}

type Visitor struct {
	gorm.Model
	Name       string
	AvatarUrl  string `gorm:"size:255;default:'';comment:头像链接"`
	GithubUser GithubUser
}

func (Visitor) TableName() string {
	return "visitors"
}

func (Visitor) GetFields() *VisitorFields {
	return &VisitorFields{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		Name:      "name",
		AvatarUrl: "avatar_url",
	}
}

func (Visitor) GetRelations() *VisitorRelations {
	return &VisitorRelations{
		GithubUser: "GithubUser",
	}
}

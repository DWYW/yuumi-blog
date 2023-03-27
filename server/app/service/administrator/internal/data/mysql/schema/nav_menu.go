package schema

import "gorm.io/gorm"

type NavMenuFields struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	ParentID    string
	Name        string
	LinkUrl     string
	ActivedRule string
	Weight      string
	Icon        string
}

type NavMenuRelations struct {
	Roles string
}

type NavMenu struct {
	gorm.Model
	ParentID    uint
	Name        string
	LinkUrl     string
	ActivedRule string
	Weight      int64
	Icon        string
	Roles       []*Role `gorm:"many2many:role_nav_menus;"`
}

func (NavMenu) TableName() string {
	return "nav_menus"
}

func (NavMenu) GetFields() *NavMenuFields {
	return &NavMenuFields{
		ID:          "id",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
		ParentID:    "parent_id",
		Name:        "name",
		LinkUrl:     "link_url",
		ActivedRule: "actived_rule",
		Weight:      "weight",
		Icon:        "icon",
	}
}

func (NavMenu) GetRelations() *NavMenuRelations {
	return &NavMenuRelations{
		Roles: "Roles",
	}
}

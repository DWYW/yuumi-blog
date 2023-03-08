package schema

import "gorm.io/gorm"

type AdministratorFields struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
	Password  string
	Salt      string
}

type Administrator struct {
	gorm.Model
	Name     string
	Password string
	Salt     string
	Roles    []*Role `gorm:"many2many:administrator_roles;"`
}

func (Administrator) GetFeilds() *AdministratorFields {
	return &AdministratorFields{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		Name:      "name",
		Password:  "password",
		Salt:      "salt",
	}
}

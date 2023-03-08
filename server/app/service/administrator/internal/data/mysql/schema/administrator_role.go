package schema

type AdministratorRoleFields struct {
	AdministratorID string
	RoleID          string
}

type AdministratorRole struct {
	AdministratorID uint
	RoleID          uint
}

func (AdministratorRole) GetFeilds() *AdministratorRoleFields {
	return &AdministratorRoleFields{
		AdministratorID: "administrator_id",
		RoleID:          "role_id",
	}
}

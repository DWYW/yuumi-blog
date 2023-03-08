package schema

type RoleNavMenuFields struct {
	RoleID    string
	NavMenuID string
}

type RoleNavMenu struct {
	RoleID    uint
	NavMenuID uint
}

func (RoleNavMenu) GetFeilds() *RoleNavMenuFields {
	return &RoleNavMenuFields{
		RoleID:    "role_id",
		NavMenuID: "nav_menu_id",
	}
}

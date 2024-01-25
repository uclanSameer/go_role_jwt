package auth

const (
	RoleAdmin = "ADMIN"
	RoleUser  = "USER"
)

type Permission struct {
	Name  string
	Roles []string
}

var Permissions = []Permission{
	{Name: "CreateUser", Roles: []string{RoleAdmin}},
	{Name: "UpdateUser", Roles: []string{RoleAdmin, RoleUser}},
	{Name: "DeleteUser", Roles: []string{RoleAdmin}},
}

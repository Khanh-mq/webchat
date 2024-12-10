package auth

type Permission struct {
	Scope    string `json:"scope"`  // pham vi su dung vi dunhuw "system"  or "room"
	Action   string `json:"action"` //  trang thai duoc dung
	Resource string `json:"resource"`
}
type Role struct {
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

var SystemRoles = map[string]Role{
	"admin": {
		Name: "admin",
		Permissions: []Permission{
			{Scope: "system", Action: "read", Resource: "*"},
			{Scope: "system", Action: "write", Resource: "*"},
		},
	},
	"group_admin": {
		Name: "group_admin",
		Permissions: []Permission{
			{Scope: "room", Action: "manage", Resource: "room"},
			{Scope: "room", Action: "read", Resource: "chat"},
		},
	},
	"member": {
		Name: "member",
		Permissions: []Permission{
			{Scope: "room", Action: "read", Resource: "chat"},
		},
	},
}

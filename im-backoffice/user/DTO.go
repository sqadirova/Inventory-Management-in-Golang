package user

type UpdateUserReq struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Enabled   bool   `json:"enabled"`
	RoleId    string `json:"role_id"`
}

type UserDTO struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Enabled   bool   `json:"enabled"`
	RoleId    string `json:"role_id"`
}

type RolesResp struct {
	Id       string `json:"id"`
	RoleType string `json:"role_type"`
}

type UserResp struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Enabled   bool      `json:"enabled"`
	Role      RolesResp `json:"role"`
}

type UserMeResp struct {
	Id        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Role      RolesResp `json:"role"`
	Username  string    `json:"username"`
}

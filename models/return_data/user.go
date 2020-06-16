package return_data

type ReturnLoginUser struct {
	Id       int    `json:"id"`
	RoleId   int    `json:"role_id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Sex      int    `json:"sex"`
	Age      int    `json:"age"`
	Avatar   string `json:"avatar"`
}

type ReturnGetUser struct {
	Id        int    `json:"id"`
	RoleId    int    `json:"role_id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Sex       int    `json:"sex"`
	Age       int    `json:"age"`
	Avatar    string `json:"avatar"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

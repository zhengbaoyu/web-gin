package models

type User struct {
	Model
	RoleId   int    `json:"role_id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Sex      int    `json:"sex"`
	Age      int    `json:"age"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
	IsDelete int    `json:"is_delete"`
}

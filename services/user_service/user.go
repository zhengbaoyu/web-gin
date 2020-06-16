package user_service

import (
	"web-gin/models"
	"web-gin/repositorys/user_repository"
)

type User struct {
	Id       int
	Phone    string
	Password string
}

func (u *User) AddUser() error {
	return user_repository.Add(models.User{
		Phone:    u.Phone,
		Password: u.Password,
	})
}
func (u *User) CheckPhone() (models.User, error) {
	return user_repository.Check(u.Phone)
}

func (u *User) GetUser() (models.User, error) {
	return user_repository.Get(u.Id)
}

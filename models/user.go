package model

import "chat-app/helper"

type User struct {
	Id       helper.UserId
	Name     string
	Email    string
	Password string
}

func (u *User) GetUser() *User {
	user := &User{
		Id:       "test_123",
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "test123",
	}
	return user
}

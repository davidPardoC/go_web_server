package main

import (
	"encoding/json"
	"net/http"
)

type MiddleWare func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type UsersList struct {
	Users []User `json:"users"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (userList *UsersList) addUser(user User) {
	userList.Users = append(userList.Users, user)
}

func (user *User) ToJson() ([]byte, error) {
	return json.Marshal(user)
}

func (users *UsersList) ToJson() ([]byte, error) {
	return json.Marshal(users)
}

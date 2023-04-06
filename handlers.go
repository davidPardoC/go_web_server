package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var usersList UsersList = UsersList{Users: []User{}}

func HandlerRoute(reswritter http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(reswritter, "Hello World")
}

func HandleHome(reswritter http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(reswritter, "this is the api endpoit")
}

func PostHandleUsers(reswritter http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(reswritter, "error: %v", err)
		return
	}
	respose, err := user.ToJson()
	usersList.addUser(user)
	if err != nil {
		reswritter.WriteHeader(http.StatusInternalServerError)
		return
	}
	reswritter.Header().Set("Content-Type", "application/json")
	reswritter.Write(respose)
}

func GetUsersHandle(reswritter http.ResponseWriter, req *http.Request) {
	respose, err := usersList.ToJson()
	if err != nil {
		reswritter.WriteHeader(http.StatusInternalServerError)
		return
	}
	reswritter.Header().Set("Content-Type", "application/json")
	reswritter.Write(respose)
}

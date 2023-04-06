package main

import "fmt"

func main() {
	server := NewServer(":8080")
	server.Handle("GET", "/", server.AddMiddleware(HandlerRoute, LogginMiddleware()))
	server.Handle("POST", "/users", PostHandleUsers)
	server.Handle("GET", "/users", GetUsersHandle)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, LogginMiddleware(), CheckAuthMiddleware()))
	err := server.Listen()
	fmt.Println(err)
}

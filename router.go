package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (router *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exists := router.rules[path]
	handler, methodExists := router.rules[path][method]
	return handler, exists, methodExists
}

func (router *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	handler, exists, methodExists := router.FindHandler(req.URL.Path, req.Method)
	if !methodExists {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if exists {
		handler(res, req)
		return
	} else {
		res.WriteHeader(http.StatusNotFound)
		return
	}

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuthMiddleware() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Checking auth")
			flag := true
			if flag {
				f(w, r)
			} else {
				return
			}

		}
	}
}

func LogginMiddleware() MiddleWare {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			handler(w, r)
		}
	}
}

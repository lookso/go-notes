package main

import (
	"context"
	"net/http"
)

type FooKey string

var UserName = FooKey("user-name")
var UserId = FooKey("user-id")

func nextHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), UserId, "100861")
		ctx2 := context.WithValue(ctx, UserName, "peanut.com")
		next(w, r.WithContext(ctx2))
	}
}

func GetUserName(context context.Context) string {
	if ret, ok := context.Value(UserName).(string); ok {
		return ret
	}
	return ""
}

func GetUserId(context context.Context) string {
	if ret, ok := context.Value(UserId).(string); ok {
		return ret
	}
	return ""
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome: "))
	w.Write([]byte(GetUserId(r.Context())))
	w.Write([]byte(" "))
	w.Write([]byte(GetUserName(r.Context())))
}

func main() {
	http.Handle("/", nextHandle(index))
	http.ListenAndServe(":8080", nil)
}



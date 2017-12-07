package users

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi"
)

func HandleGetUsersWithUserId() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("GET /users/{userId}: user id is %v", chi.URLParam(r, "userId"))))
	}
}

func HandleGetUsers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("userId")
		w.Write([]byte(fmt.Sprintf("GET /users: user id param is %v", userId)))
	}
}
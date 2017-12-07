package users

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi"
)

func HandleGetUsers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("GET /users/{userId}: view user id %v", chi.URLParam(r, "userId"))))
	}
}

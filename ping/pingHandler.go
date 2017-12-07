package ping

import "net/http"

func HandleGetPing() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("the-header")
		w.Write([]byte("pong.  you sent " + h))
	}
}
package controller

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gdichter/rest-controller-example/dtos"
)

func MyRouter() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/ping", handleGetPing())

	r.Route("/stuff", func(r chi.Router) {
		r.Get("/", handleGetStuff())

		r.Post("/", handlePostStuff())
	})
	http.ListenAndServe(":3000", r)

}
func handleGetPing() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("the-header")
		w.Write([]byte("pong.  you sent " + h))
	}
}
func handleGetStuff() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from Get: here is some stuff"))
	}
}

func handlePostStuff() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		response := handleRequest(body)
		w.Write(response)
	}
}

func handleRequest(bytes []byte) []byte {
	incomingRequest := request.MyRequest{}
	unmarshallErr := json.Unmarshal(bytes, &incomingRequest)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}

	originalArg1 := incomingRequest.Arg1
	incomingRequest.Arg1 = originalArg1 + "-NEW"

	originalA1 := incomingRequest.Arr[0].A1
	incomingRequest.Arr[0].A1 = originalA1 + 1

	originalB1 := incomingRequest.Arr[0].B1
	incomingRequest.Arr[0].B1 = originalB1 + "-UPDATED"

	data, err := json.Marshal(&incomingRequest)
	if err != nil {
		panic(err)
	}
	return data


}

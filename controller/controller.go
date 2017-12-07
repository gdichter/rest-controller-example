package controller

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
	"github.com/gdichter/rest-controller-example/stuff"
	"github.com/gdichter/rest-controller-example/ping"
)

func MyRouter() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/ping", ping.HandleGetPing())

	r.Route("/stuff", func(r chi.Router) {
		r.Get("/", stuff.HandleGetStuff())

		r.Post("/", stuff.HandlePostStuff())
	})
	http.ListenAndServe(":3000", r)

}



package main

import (
	"net/http"

	"project_name/application/controller"
	"project_name/application/interface/handler"
	v1 "project_name/application/interface/handler/v1"
	"project_name/application/interface/middleware"
	"project_name/application/pkg/router"

	"github.com/go-chi/chi"
	"github.com/urfave/negroni"
)

func buildRouter() http.Handler {
	recovery := negroni.NewRecovery()
	recovery.Formatter = &middleware.PanicFormatter{}
	recovery.PanicHandlerFunc = middleware.Logging

	n := negroni.New()
	n.Use(recovery)

	chiContext := router.NewChiContext()

	var router = chi.NewRouter()
	router.Get("/healthz", handler.Healthz)
	router.Route("/v1", func(r chi.Router) {
		r.Route("/", v1Router(chiContext))
	})
	n.UseHandler(router)
	return n
}

func v1Router(chiContext *router.ChiContext) func(chi.Router) {
	var api = v1.NewHello(chiContext, controller.NewHello())

	return func(r chi.Router) {
		r.Route("/hello", func(r chi.Router) {
			r.Get("/", api.Get)
		})
	}
}

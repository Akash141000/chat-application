package router

import (
	"chat-app/controller"
	"chat-app/helper"
	"chat-app/middleware"
	"fmt"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

type Router struct {
	*http.ServeMux
}

func New(mux *http.ServeMux) *Router {
	r := &Router{
		ServeMux: mux,
	}
	r.RegisterRoutes()
	return r
}

func (r *Router) RegisterRoutes() {
	//
	r.Route("/login", controller.Login, POST, false)
	r.Route("/signup", controller.SignUp, POST, false)

	//
	r.Route("/chat", controller.Chat, POST, true)

	//default route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from chat-app!")
	})
}

func (r *Router) Route(pattern string, controller helper.ControllerFunc, method string, auth bool) {
	if auth {
		r.HandleFunc(pattern, middleware.Logger(middleware.AuthHandler(middleware.ErrorHandler(controller, method))))
	} else {
		r.HandleFunc(pattern, middleware.Logger(middleware.ErrorHandler(controller, method)))
	}
}

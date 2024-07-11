package http

import (
	"github.com/brilianpmw/synapsis/internal/pkg/middleware"
	"github.com/brilianpmw/synapsis/presentation"
	"github.com/go-chi/chi"
)

type HttpHandler struct {
	router  *chi.Mux
	usecase presentation.IUserUC
}

func New(router *chi.Mux, usecase presentation.IUserUC) *HttpHandler {
	return &HttpHandler{
		router:  router,
		usecase: usecase,
	}
}

func (handler *HttpHandler) SetEndpoint() {
	// Internal API
	handler.router.Post("/login", handler.Login)
	handler.router.Get("/user", middleware.Authenticate(handler.Welcome))

}

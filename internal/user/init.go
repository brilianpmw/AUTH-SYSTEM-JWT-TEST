package missionmanagement

import (
	DelivUserHTTP "github.com/brilianpmw/linknau/internal/user/delivery/http"
	"github.com/brilianpmw/linknau/internal/user/usecase"
	"github.com/brilianpmw/linknau/presentation"
	"github.com/go-chi/chi"
)

func NewHTTP(
	router *chi.Mux,
	user presentation.IUser,
) (err error) {

	uc := usecase.New(&usecase.Repositories{
		User: user,
	})
	h := DelivUserHTTP.New(router, uc)
	h.SetEndpoint()
	return
}

package main

import (
	"fmt"
	"net/http"

	cfg "github.com/brilianpmw/linknau/internal/pkg/config"
	pqRepo "github.com/brilianpmw/linknau/internal/repository/postgre"
	httpUser "github.com/brilianpmw/linknau/internal/user/delivery/http"
	"github.com/brilianpmw/linknau/internal/user/usecase"
	"github.com/go-chi/chi"
)

var postgre *pqRepo.Postgre

func main() {
	cfg.NewConfig()
	postgre, _ := pqRepo.New()
	authUseCase := usecase.New(&usecase.Repositories{
		User: postgre,
	})
	router := chi.NewRouter()

	deliveryUserHTTP := httpUser.New(router, authUseCase)
	deliveryUserHTTP.SetEndpoint()

	fmt.Printf("RUNNING at port :8080")

	http.ListenAndServe(":8080", router)

}

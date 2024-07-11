package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brilianpmw/synapsis/presentation"
)

func (a *HttpHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds presentation.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := a.usecase.DoLogin(context.Background(), creds)
	fmt.Printf("ERROR :%v,", err)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (a *HttpHandler) Welcome(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome "})
}

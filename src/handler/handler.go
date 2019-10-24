package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"src/src/service"

	"github.com/gorilla/mux"
)

type Funcs struct {
	Ctx    context.Context
	Savage service.Savage
}

func (h *Funcs) GetSavage(w http.ResponseWriter, r *http.Request) {

	requestParameters := mux.Vars(r)
	number := requestParameters["number"]

	savage, err := h.Savage.Get(number)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(savage)
}

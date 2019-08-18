package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"src/src/service"
)

type Funcs struct {
	Ctx    context.Context
	Savage service.Savage
}

func (h *Funcs) GetSavage(w http.ResponseWriter, r *http.Request) {
	savage, err := h.Savage.Get()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(savage)
}

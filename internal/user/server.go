package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewUserHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

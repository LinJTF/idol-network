package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service Service
}

func NewUserHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting users...")
	ctx := r.Context()

	users, err := h.service.GetUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting user by ID...")
	ctx := r.Context()

	userID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.service.GetUserByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting user by email...")
	ctx := r.Context()

	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Missing email parameter", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByEmail(ctx, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Creating user...")
	ctx := r.Context()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser, err := h.service.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)

}

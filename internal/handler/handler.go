package handler

import (
	"context"
	"encoding/json"
	"errors"
	"go-unit-of-work-example/internal/entity"
	"net/http"
)

type Service interface {
	RegisterNewUser(ctx context.Context, name string, email string) error
}

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

type Handler struct {
	service Service
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.service.RegisterNewUser(r.Context(), r.URL.Query().Get("name"), r.URL.Query().Get("email"))
	if err != nil {
		if errors.Is(err, entity.ErrAlreadyRegistered) {
			http.Error(w, entity.ErrAlreadyRegistered.Error(), http.StatusConflict)
			return
		}

		// TODO log it
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

package v1

import (
	"encoding/json"
	"net/http"
	"user-management-api/transport/api/v1/model"

	"github.com/go-chi/chi"
)

func getEmailFromURL(r *http.Request) string {
	email := chi.URLParam(r, "email")
	return email
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Decode JSON
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Validate birthday
	if user.BirthDay.Validate() != nil {
		http.Error(w, "Invalid birthDay", http.StatusBadRequest)
		return
	}

	// Call service
	h.service.CreateUser(r.Context(), user)

	// Respond OK
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented - GetUser")
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	panic("not implemented - ListUsers")
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented - UpdateUser")
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented - DeleteUser")
}

package handler

import (
	"encoding/json"
	"net/http"

	"achievement-leveling/achievement/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) UpsertUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		GoogleID string `json:"google_id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Picture  string `json:"picture"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	user, err := h.svc.UpsertFromGoogle(req.GoogleID, req.Email, req.Name, req.Picture)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")
	user, err := h.svc.GetByID(userID)
	if err != nil {
		http.Error(w, `{"error":"user not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

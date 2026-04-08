package handler

import (
	"encoding/json"
	"net/http"

	"achievement-leveling/achievement/service"
)

type ProgressHandler struct {
	svc *service.LevelingService
}

func NewProgressHandler(svc *service.LevelingService) *ProgressHandler {
	return &ProgressHandler{svc: svc}
}

func (h *ProgressHandler) CompleteAchievement(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	var req struct {
		AchievementID string `json:"achievement_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	result, err := h.svc.CompleteAchievement(userID, req.AchievementID)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *ProgressHandler) GetUserProgress(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	progress, err := h.svc.GetUserProgress(userID)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(progress)
}

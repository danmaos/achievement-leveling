package handler

import (
	"encoding/json"
	"net/http"

	"achievement-leveling/achievement/service"
)

type AchievementHandler struct {
	svc *service.AchievementService
}

func NewAchievementHandler(svc *service.AchievementService) *AchievementHandler {
	return &AchievementHandler{svc: svc}
}

func (h *AchievementHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categoryID := r.URL.Query().Get("category_id")
	var achievements interface{}
	var err error

	if categoryID != "" {
		achievements, err = h.svc.GetByCategory(categoryID)
	} else {
		achievements, err = h.svc.GetAll()
	}

	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(achievements)
}

func (h *AchievementHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	a, err := h.svc.GetByID(id)
	if err != nil {
		http.Error(w, `{"error":"achievement not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

func (h *AchievementHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.svc.GetAllCategories()
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

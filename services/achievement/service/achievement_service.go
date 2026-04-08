package service

import (
	"achievement-leveling/achievement/model"
	"achievement-leveling/achievement/repository"
)

type AchievementService struct {
	repo *repository.AchievementRepo
}

func NewAchievementService(repo *repository.AchievementRepo) *AchievementService {
	return &AchievementService{repo: repo}
}

func (s *AchievementService) GetAll() ([]model.Achievement, error) {
	return s.repo.FindAll()
}

func (s *AchievementService) GetByID(id string) (*model.Achievement, error) {
	return s.repo.FindByID(id)
}

func (s *AchievementService) GetByCategory(categoryID string) ([]model.Achievement, error) {
	return s.repo.FindByCategory(categoryID)
}

func (s *AchievementService) GetAllCategories() ([]model.AchievementCategory, error) {
	return s.repo.FindAllCategories()
}

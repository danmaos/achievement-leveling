package repository

import (
	"achievement-leveling/achievement/model"

	"gorm.io/gorm"
)

type AchievementRepo struct {
	db *gorm.DB
}

func NewAchievementRepo(db *gorm.DB) *AchievementRepo {
	return &AchievementRepo{db: db}
}

func (r *AchievementRepo) FindAll() ([]model.Achievement, error) {
	var achievements []model.Achievement
	err := r.db.Preload("Category").Find(&achievements).Error
	return achievements, err
}

func (r *AchievementRepo) FindByID(id string) (*model.Achievement, error) {
	var a model.Achievement
	err := r.db.Preload("Category").First(&a, "id = ?", id).Error
	return &a, err
}

func (r *AchievementRepo) FindByCategory(categoryID string) ([]model.Achievement, error) {
	var achievements []model.Achievement
	err := r.db.Preload("Category").Where("category_id = ?", categoryID).Find(&achievements).Error
	return achievements, err
}

func (r *AchievementRepo) FindAllCategories() ([]model.AchievementCategory, error) {
	var categories []model.AchievementCategory
	err := r.db.Find(&categories).Error
	return categories, err
}

package repository

import (
	"achievement-leveling/achievement/model"

	"gorm.io/gorm"
)

type ProgressRepo struct {
	db *gorm.DB
}

func NewProgressRepo(db *gorm.DB) *ProgressRepo {
	return &ProgressRepo{db: db}
}

func (r *ProgressRepo) FindByUser(userID string) ([]model.UserAchievement, error) {
	var progress []model.UserAchievement
	err := r.db.Preload("Achievement.Category").Where("user_id = ?", userID).Find(&progress).Error
	return progress, err
}

func (r *ProgressRepo) FindByUserAndAchievement(userID, achievementID string) (*model.UserAchievement, error) {
	var ua model.UserAchievement
	err := r.db.Where("user_id = ? AND achievement_id = ?", userID, achievementID).First(&ua).Error
	return &ua, err
}

func (r *ProgressRepo) Create(ua *model.UserAchievement) error {
	return r.db.Create(ua).Error
}

func (r *ProgressRepo) Update(ua *model.UserAchievement) error {
	return r.db.Save(ua).Error
}

func (r *ProgressRepo) GetLevelForXP(xp int) (*model.LevelThreshold, error) {
	var threshold model.LevelThreshold
	err := r.db.Where("xp_required <= ?", xp).Order("level DESC").First(&threshold).Error
	return &threshold, err
}

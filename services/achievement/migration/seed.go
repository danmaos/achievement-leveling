package migration

import (
	"log"

	"achievement-leveling/achievement/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.AchievementCategory{},
		&model.Achievement{},
		&model.UserAchievement{},
		&model.LevelThreshold{},
	)
	if err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}
}

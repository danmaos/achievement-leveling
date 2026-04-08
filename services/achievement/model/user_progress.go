package model

import "time"

type UserAchievement struct {
	ID             string      `json:"id" gorm:"primaryKey;type:char(36)"`
	UserID         string      `json:"user_id" gorm:"type:char(36);not null"`
	AchievementID  string      `json:"achievement_id" gorm:"type:char(36);not null"`
	Achievement    Achievement `json:"achievement,omitempty" gorm:"foreignKey:AchievementID"`
	CompletedAt    time.Time   `json:"completed_at"`
	TimesCompleted int         `json:"times_completed" gorm:"default:1"`
}

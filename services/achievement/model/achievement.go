package model

import "time"

type AchievementCategory struct {
	ID        string    `json:"id" gorm:"primaryKey;type:char(36)"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Icon      string    `json:"icon" gorm:"type:varchar(100)"`
	Color     string    `json:"color" gorm:"type:varchar(7)"`
	CreatedAt time.Time `json:"created_at"`
}

type Achievement struct {
	ID          string              `json:"id" gorm:"primaryKey;type:char(36)"`
	CategoryID  string              `json:"category_id" gorm:"type:char(36)"`
	Category    AchievementCategory `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Title       string              `json:"title" gorm:"type:varchar(255);not null"`
	Description string              `json:"description" gorm:"type:text"`
	XPReward    int                 `json:"xp_reward" gorm:"default:10;not null"`
	Icon        string              `json:"icon" gorm:"type:varchar(100)"`
	Difficulty  string              `json:"difficulty" gorm:"type:enum('easy','medium','hard','epic');default:'medium'"`
	Repeatable  bool                `json:"repeatable" gorm:"default:false"`
	CreatedAt   time.Time           `json:"created_at"`
}

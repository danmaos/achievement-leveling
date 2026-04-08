package model

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:char(36)"`
	GoogleID  string    `json:"google_id" gorm:"type:varchar(255);uniqueIndex;not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Picture   string    `json:"picture" gorm:"type:varchar(512)"`
	XP        int       `json:"xp" gorm:"default:0;not null"`
	Level     int       `json:"level" gorm:"default:1;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

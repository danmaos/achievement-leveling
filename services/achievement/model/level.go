package model

type LevelThreshold struct {
	Level      int    `json:"level" gorm:"primaryKey"`
	XPRequired int    `json:"xp_required" gorm:"not null"`
	Title      string `json:"title" gorm:"type:varchar(100)"`
}

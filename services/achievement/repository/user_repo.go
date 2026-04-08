package repository

import (
	"achievement-leveling/achievement/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Upsert(user *model.User) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "google_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"email", "name", "picture", "updated_at"}),
	}).Create(user).Error
}

func (r *UserRepo) FindByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}

func (r *UserRepo) FindByGoogleID(googleID string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "google_id = ?", googleID).Error
	return &user, err
}

func (r *UserRepo) Update(user *model.User) error {
	return r.db.Save(user).Error
}

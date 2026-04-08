package service

import (
	"achievement-leveling/achievement/model"
	"achievement-leveling/achievement/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UpsertFromGoogle(googleID, email, name, picture string) (*model.User, error) {
	existing, err := s.repo.FindByGoogleID(googleID)
	if err == nil {
		existing.Email = email
		existing.Name = name
		existing.Picture = picture
		if err := s.repo.Update(existing); err != nil {
			return nil, err
		}
		return existing, nil
	}

	user := &model.User{
		ID:       uuid.New().String(),
		GoogleID: googleID,
		Email:    email,
		Name:     name,
		Picture:  picture,
		XP:       0,
		Level:    1,
	}
	if err := s.repo.Upsert(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetByID(id string) (*model.User, error) {
	return s.repo.FindByID(id)
}

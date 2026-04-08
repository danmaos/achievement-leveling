package service

import (
	"errors"
	"time"

	"achievement-leveling/achievement/model"
	"achievement-leveling/achievement/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LevelingService struct {
	progressRepo *repository.ProgressRepo
	userRepo     *repository.UserRepo
	achRepo      *repository.AchievementRepo
}

func NewLevelingService(pr *repository.ProgressRepo, ur *repository.UserRepo, ar *repository.AchievementRepo) *LevelingService {
	return &LevelingService{progressRepo: pr, userRepo: ur, achRepo: ar}
}

type CompleteResult struct {
	User        *model.User `json:"user"`
	Achievement *model.Achievement `json:"achievement"`
	XPGained    int `json:"xp_gained"`
	LeveledUp   bool `json:"leveled_up"`
	NewLevel    *model.LevelThreshold `json:"new_level,omitempty"`
}

func (s *LevelingService) CompleteAchievement(userID, achievementID string) (*CompleteResult, error) {
	achievement, err := s.achRepo.FindByID(achievementID)
	if err != nil {
		return nil, errors.New("achievement not found")
	}

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	existing, err := s.progressRepo.FindByUserAndAchievement(userID, achievementID)
	if err == nil && !achievement.Repeatable {
		return nil, errors.New("achievement already completed and is not repeatable")
	}

	if err == nil && achievement.Repeatable {
		existing.TimesCompleted++
		existing.CompletedAt = time.Now()
		if err := s.progressRepo.Update(existing); err != nil {
			return nil, err
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		ua := &model.UserAchievement{
			ID:             uuid.New().String(),
			UserID:         userID,
			AchievementID:  achievementID,
			CompletedAt:    time.Now(),
			TimesCompleted: 1,
		}
		if err := s.progressRepo.Create(ua); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	oldLevel := user.Level
	user.XP += achievement.XPReward

	threshold, err := s.progressRepo.GetLevelForXP(user.XP)
	if err == nil {
		user.Level = threshold.Level
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	result := &CompleteResult{
		User:        user,
		Achievement: achievement,
		XPGained:    achievement.XPReward,
		LeveledUp:   user.Level > oldLevel,
	}
	if result.LeveledUp {
		result.NewLevel = threshold
	}

	return result, nil
}

func (s *LevelingService) GetUserProgress(userID string) ([]model.UserAchievement, error) {
	return s.progressRepo.FindByUser(userID)
}

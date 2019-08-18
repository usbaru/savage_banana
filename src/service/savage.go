package service

import (
	"src/src/model"
	"src/src/repository"
)

type Savage interface {
	Get() (*model.Savage, error)
}

type savageRepository struct {
	savage repository.Savage
}

func NewSavageService(savage repository.Savage) *savageRepository {
	return &savageRepository{
		savage: savage,
	}
}

func (s *savageRepository) Get() (*model.Savage, error) {

	savage := model.Savage{
		UserID:    1,
		ID:        1,
		Title:     "Savage",
		Completed: true,
	}
	return &savage, nil
}

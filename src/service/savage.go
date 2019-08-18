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

	savage, err := s.savage.GetSavage()
	if err != nil {
		return nil, err
	}
	return &savage, nil
}

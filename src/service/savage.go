package service

import (
	"src/src/model"
	"src/src/repository"
)

type Savage interface {
	Get(number string) (*model.Savage, error)
}

type savageRepository struct {
	savage repository.Savage
}

func NewSavageService(savage repository.Savage) *savageRepository {
	return &savageRepository{
		savage: savage,
	}
}

func (s *savageRepository) Get(number string) (*model.Savage, error) {

	savage, err := s.savage.GetSavage(number)
	if err != nil {
		return nil, err
	}

	return savage, nil
}

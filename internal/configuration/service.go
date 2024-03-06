package configuration

import (
	"errors"
	"fmt"

	"github.com/unm4sked/finch/internal/entities"
)

type Service interface {
	Create() error
	GetConfigurations() ([]entities.Configuration, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create() error {
	s.repository.CreateConfiguration()
	return errors.New("not implemented")
}

func (s *service) GetConfigurations() ([]entities.Configuration, error) {
	configs, err := s.repository.GetConfigurations()
	var configsDefault []entities.Configuration
	if err != nil {
		fmt.Println("Error Service here")
		return configsDefault, err
	}
	return configs, nil
}

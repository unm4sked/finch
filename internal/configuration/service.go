package configuration

import (
	"errors"
)

type Service interface {
	Create() error
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
	return errors.New("error")
}

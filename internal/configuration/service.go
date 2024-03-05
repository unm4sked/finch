package configuration

import "errors"

type Service interface {
	Create() error
}

type service struct {
	// Deps
}

func New() Service {
	return &service{}
}

func (s *service) Create() error {
	return errors.New("error")
}

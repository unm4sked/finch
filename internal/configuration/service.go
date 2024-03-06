package configuration

import (
	"fmt"

	"github.com/unm4sked/finch/internal/entities"
)

type Service interface {
	Create(description string) (string, error)
	GetConfigurations() ([]entities.Configuration, error)
	GetConfiguration(id string) (entities.Configuration, error)
	RemoveConfiguration(id string) error
	UpdateConfigurationDescription(id string, description string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(description string) (string, error) {
	id, err := s.repository.CreateConfiguration(description)
	if err != nil {
		return "", err
	}
	return id, nil
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
func (s *service) GetConfiguration(id string) (entities.Configuration, error) {
	config, err := s.repository.GetConfigurationById(id)
	var configDefault entities.Configuration

	if err != nil {
		fmt.Println("Error Service here")
		return configDefault, err
	}
	return config, nil
}

func (s *service) RemoveConfiguration(id string) error {
	err := s.repository.DeleteConfiguration(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateConfigurationDescription(id string, description string) error {
	err := s.repository.UpdateConfiguration(id, description)
	if err != nil {
		return err
	}
	return nil
}

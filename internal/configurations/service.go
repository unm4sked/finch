package configurations

import (
	"fmt"

	"github.com/google/uuid"
)

type Service interface {
	Create(description string) (string, error)
	GetConfigurations() ([]Configuration, error)
	GetConfiguration(id string) (Configuration, error)
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
	id := uuid.New().String()

	err := s.repository.CreateConfiguration(id, description)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *service) GetConfigurations() ([]Configuration, error) {
	configs, err := s.repository.GetConfigurations()
	var configsDefault []Configuration
	if err != nil {
		fmt.Println("Error Service here")
		return configsDefault, err
	}
	return configs, nil
}
func (s *service) GetConfiguration(id string) (Configuration, error) {
	config, err := s.repository.GetConfigurationById(id)
	var configDefault Configuration
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

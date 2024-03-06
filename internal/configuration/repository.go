package configuration

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/unm4sked/finch/pkg/postgres"
)

type Repository interface {
	CreateConfiguration() error
	GetConfigurationById() error
	GetConfigurations() error
	DeleteConfiguration() error
	UpdateConfiguration() error
}

type repository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db postgres.Postgres) Repository {
	return &repository{
		db: db.Pool,
	}
}

func (r *repository) CreateConfiguration() error {
	description := "text text text"
	tag, err := r.db.Exec(context.Background(), `INSERT INTO configurations (description) VALUES ($1)`, description)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println(tag)
	return errors.New("hello")
}

func (r *repository) GetConfigurationById() error {
	row := r.db.QueryRow(context.Background(), "select * from configurations LIMIT 1")
	err := row.Scan()
	if err != nil {
		log.Println("Error while getting configuration")
		return err
	}
	return nil
}

func (r *repository) GetConfigurations() error {
	return errors.New("not implemented")
}

func (r *repository) DeleteConfiguration() error {
	return errors.New("not implemented")
}

func (r *repository) UpdateConfiguration() error {
	return errors.New("not implemented")
}

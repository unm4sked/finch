package resources

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/unm4sked/finch/pkg/postgres"
)

type ResourceID string

type Repository interface {
	CreateResource(user Resource) (ResourceID, error)
	GetResource(id ResourceID) (Resource, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db postgres.Postgres) Repository {
	return &repository{
		db: db.Pool,
	}
}

func (u *repository) CreateResource(user Resource) (ResourceID, error) {
	return "TBD", nil
}

func (u *repository) GetResource(id ResourceID) (Resource, error) {
	return Resource{}, nil
}

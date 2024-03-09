package users

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/unm4sked/finch/pkg/postgres"
)

type UserID string

type Repository interface {
	CreateUser(user User) (UserID, error)
	GetUser(id UserID) (User, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db postgres.Postgres) Repository {
	return &repository{
		db: db.Pool,
	}
}

func (u *repository) CreateUser(user User) (UserID, error) {
	return "TBD", nil
}

func (u *repository) GetUser(id UserID) (User, error) {
	return User{}, nil
}

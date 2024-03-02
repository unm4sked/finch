package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New() (*Postgres, error) {
	connStr := "user=postgres dbname=app sslmode=disable password=password host=localhost"
	pgConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config to connect to database: %v\n", err)
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		return nil, err
	}

	return &Postgres{
		Pool: pool,
	}, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
	fmt.Println("Database connection closed")
}

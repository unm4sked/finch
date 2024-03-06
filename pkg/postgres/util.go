package postgres

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize int
	connTimeout time.Duration
	Pool        *pgxpool.Pool
}

func Connect(opts ...Option) (*Postgres, error) {

	pg := &Postgres{
		maxPoolSize: _defaultMaxPoolSize,
		connTimeout: _defaultConnAttempts,
	}

	for _, opt := range opts {
		opt(pg)
	}

	connStr := "user=postgres dbname=app sslmode=disable password=password host=localhost port=5432"
	pgConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config to connect to database: %v\n", err)
		return nil, err
	}

	pgConfig.MaxConns = int32(pg.maxPoolSize)
	pgConfig.MaxConns = int32(pg.maxPoolSize)

	pg.Pool, err = pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	if err := pg.Pool.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		return nil, err
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
	fmt.Println("Database connection closed")
}

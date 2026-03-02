package postgres

import (
	"context"
	"fmt"

	"github.com/EduCoelhoTs/base-hex-arq-api/internal/infra/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Db *pgxpool.Pool
}

func NewPostgres(ctx context.Context, config *config.Config) (*Postgres, error) {

	connstr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	conn, err := pgxpool.New(ctx, connstr)
	if err != nil {
		return nil, fmt.Errorf("error to connect to database: %w", err)
	}

	return &Postgres{Db: conn}, nil
}

func (p *Postgres) Close() {
	p.Db.Close()
}

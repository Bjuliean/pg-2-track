package postgres

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	logger *slog.Logger
	config *Config
	db     *pgxpool.Pool
}

func NewDB(config *Config, logger *slog.Logger) *Service {
	nwdb, err := pgxpool.New(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.User, config.Password, config.Host, config.Port, config.Name))
	if err != nil {
		log.Fatal(err)
	}

	return &Service{
		logger: logger,
		config: config,
		db: nwdb,
	}
}

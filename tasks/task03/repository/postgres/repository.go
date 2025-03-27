package postgres

import (
	"log/slog"
	"pg-2-track/tasks/task03/repository"
)

type repo struct {
	*Service
	logger *slog.Logger
}

func NewRepository(db *Service, logger *slog.Logger) repository.Repository {
	return &repo{
		Service: db,
		logger: logger,
	}
}
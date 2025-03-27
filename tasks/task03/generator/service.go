package generator

import (
	"log/slog"
	"pg-2-track/tasks/task03/repository"
)

type Service struct {
	repo repository.Repository
	config *Config
	logger *slog.Logger
}

func New(config *Config, repo repository.Repository, logger *slog.Logger) *Service {
	return &Service{
		repo: repo,
		config: config,
		logger: logger,
	}
}
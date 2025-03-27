package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"pg-2-track/tasks/task03/config"
	"pg-2-track/tasks/task03/generator"
	"pg-2-track/tasks/task03/repository/postgres"
)

type Config struct {
	DB        postgres.Config  `envPrefix:"PG_"`
	Generator generator.Config `envPrefix:"GEN_"`
}

func main() {
	ctx := context.Background()

	cfg := Config{}

	err := config.ReadConfig("none", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	logs := &slog.Logger{}

	db := postgres.NewDB(&cfg.DB, logs)
	repo := postgres.NewRepository(db, logs)
	gen := generator.New(&cfg.Generator, repo, logs)

	fmt.Println("GENERATION STARTED")

	err = gen.Generate(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GENERATED %d OBJECTS", cfg.Generator.NumberOfGenerations)
}

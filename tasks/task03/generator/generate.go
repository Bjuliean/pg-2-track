package generator

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"pg-2-track/tasks/task03/repository"

	"golang.org/x/sync/errgroup"
)

func (s *Service) Generate(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(s.config.NumberOfGoroutines)

	generationsLeft := s.config.NumberOfGenerations

	generationProccess := 0

	for generationsLeft > 0 {
		generationsLeft -= s.config.BatchSize
		curBatch := s.config.BatchSize
		if generationsLeft < 0 {
			curBatch += generationsLeft
		}

		g.Go(func() error {
			err := s.generationIteration(ctx, curBatch)
			fmt.Printf("GENERATION PROCCESS %d FINISHED\n", generationProccess)
			return err
		})

		generationProccess++
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (s *Service) generationIteration(ctx context.Context, totalGenerations int) error {
	nwcars := s.createCars(totalGenerations)

	err := s.repo.SaveCars(ctx, nwcars)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Service) createCars(numberOfCars int) []repository.Car {
	result := make([]repository.Car, 0, numberOfCars)

	for range numberOfCars {
		newCar := repository.Car{
			Brand: generateWord(
				rand.Intn(s.config.MaxCarBrandLen-s.config.MinCarBrandLen) + s.config.MinCarBrandLen,
			),
			Price: int64(generatePrice(s.config.MinCarPrice, s.config.MaxCarPrice)),
		}

		result = append(result, newCar)
	}

	return result
}

package postgres

import (
	"context"
	"fmt"
	"pg-2-track/tasks/task03/repository"

	"github.com/jackc/pgx/v5"
)

func (r *repo) SaveCars(ctx context.Context, cars []repository.Car) error {
	if len(cars) == 0 {
		return nil
	}

	conn, err := r.db.Acquire(ctx)
	if err != nil {
		fmt.Println(err.Error())

		return err
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	for _, car := range cars {
		batch.Queue(`
		INSERT INTO cars(brand, price)
		VALUES (@Brand, @Price)`,
			pgx.NamedArgs{
				"Brand": car.Brand,
				"Price": car.Price,
			})
	}

	batchResult := conn.SendBatch(ctx, batch)

	for range len(cars) {
		_, err := batchResult.Exec()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

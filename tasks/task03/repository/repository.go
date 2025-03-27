package repository

import "context"

type Repository interface {
	SaveCars(ctx context.Context, cars []Car) error
}

type Car struct {
	Brand string `db:"brand"`
	Price int64  `db:"price"`
}

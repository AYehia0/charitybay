// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"charitybay/internal/entity"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// UseCase -.
	Greet interface {
		CreateGreet(context.Context, entity.Greet) error
		GetGreets(context.Context) ([]entity.Greet, error)
	}

	// GreetRepo -.
	GreetRepo interface {
		CreateGreet(context.Context, entity.Greet) error
		GetGreets(context.Context) ([]entity.Greet, error)
	}
)

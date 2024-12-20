package usecase

import (
	"charitybay/internal/entity"
	"context"
)

type GreetUsecase struct {
	repo GreetRepo
}

func New(r GreetRepo) *GreetUsecase {
	return &GreetUsecase{
		repo: r,
	}
}

func (uc *GreetUsecase) GetGreets(ctx context.Context) ([]entity.Greet, error) {
	greets, err := uc.repo.GetGreets(ctx)
	if err != nil {
		return nil, err
	}
	return greets, nil
}

func (uc *GreetUsecase) CreateGreet(ctx context.Context, greet entity.Greet) error {
	err := uc.repo.CreateGreet(ctx, greet)
	if err != nil {
		return err
	}
	return nil
}

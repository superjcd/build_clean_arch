package usecase

import (
	"context"
	"greattime_coffee/internal/entity"

	"fmt"
)

type CoffeeServiceUsecase struct {
	repo   CoffeeRepo
	cuppon CouponService
}

func New(r CoffeeRepo, c CouponService) *CoffeeServiceUsecase {
	return &CoffeeServiceUsecase{
		repo:   r,
		cuppon: c,
	}
}

func (uc CoffeeServiceUsecase) List(ctx context.Context) ([]entity.Coffee, error) {
	coffees, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("GetAll coffees from Repo fail: %w", err)
	}
	return coffees, nil
}

func (uc CoffeeServiceUsecase) GetPrice(ctx context.Context, name string) (float64, error) {
	coffee, err := uc.repo.GetOne(ctx, name)

	if err != nil {
		return 0, err
	}

	coupon, err := uc.cuppon.GetCoupon(ctx)

	if err != nil {
		return 0, err
	}

	return (coffee.Price - coupon), nil
}

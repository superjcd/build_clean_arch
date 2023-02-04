package usecase

import (
	"context"
	"greattime_coffee/internal/entity"
)

type (
	CoffeeService interface {
		List(context.Context) ([]entity.Coffee, error)
		GetPrice(context.Context, string) (float64, error)
	}

	CoffeeRepo interface {
		GetAll(context.Context) ([]entity.Coffee, error)
		GetOne(context.Context, string) (entity.Coffee, error)
	}

	CouponService interface {
		GetCoupon(context.Context) (float64, error)
	}
)

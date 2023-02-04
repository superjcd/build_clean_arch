package repo

import (
	"context"
	"greattime_coffee/internal/entity"
)

type InMemoryRepo struct {
	data []entity.Coffee
}

func NewInMemoryRepo() *InMemoryRepo {
	coffees := []entity.Coffee{
		{1, "Cappucino", "Cappucino", 34},
		{2, "Caffè Americano", "Americano", 28},
		{3, "Caffè Misto", "Brewed", 30},
		{4, "Pistachio Latte", "Latte", 30},
	}
	return &InMemoryRepo{data: coffees}
}

func (repo *InMemoryRepo) GetAll(ctx context.Context) ([]entity.Coffee, error) {
	return repo.data, nil

}

func (repo *InMemoryRepo) GetOne(ctx context.Context, name string) (entity.Coffee, error) {
	var Target entity.Coffee

	for _, coffee := range repo.data {
		// Find the first coffee
		if coffee.Name == name {
			Target = coffee
			break
		}
	}

	return Target, nil
}

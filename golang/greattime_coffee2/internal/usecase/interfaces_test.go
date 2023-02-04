package usecase

import (
	"context"
	"greattime_coffee/internal/entity"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestGetCoffePrice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test_coffee := entity.Coffee{
		ID:    1,
		Name:  "test_coffee",
		Type:  "test_type",
		Price: 25.5,
	}

	// get coupon
	m_c := NewMockCouponService(ctrl)
	m_c.EXPECT().GetCoupon(gomock.Any()).Return(10.0, nil)

	//
	m_r := NewMockCoffeeRepo(ctrl)
	m_r.EXPECT().GetOne(gomock.Any(), gomock.Eq("test_coffee")).Return(test_coffee, nil)

	new_service := New(m_r, m_c)

	if p, err :=new_service.GetPrice(context.Background(),"test_coffee"); err != nil {
		t.Fatal(err)
	}else {
		if p != 15.5 {
			t.Fatal("Price error")
		}
	}

}

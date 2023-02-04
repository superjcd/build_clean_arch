package coupon

import (
	"context"
	"math/rand"
	"time"
)

type RandomCoupon struct {
	seed int64
}

func NewRandomCuppon() *RandomCoupon {
	return &RandomCoupon{seed: time.Now().UnixNano()}
}

func (coup *RandomCoupon) GetCoupon(ctx context.Context) (float64, error) {
	rand.Seed(coup.seed)
	min := 0
	max := 10
	coupon := rand.Intn(max-min+1) + min

	return float64(coupon), nil
}

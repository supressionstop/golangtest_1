package sport

import (
	"time"
)

type Baseball struct {
	rate      float64 // todo
	createdAt time.Time
}

func NewBaseball(rate float64, createdAt time.Time) *Baseball { // todo rate
	return &Baseball{
		rate:      rate,
		createdAt: createdAt,
	}
}

func (s Baseball) Name() string {
	return "baseball"
}

func (s Baseball) Rate() float64 { // todo
	return s.rate
}

func (s Baseball) CreatedAt() time.Time {
	return s.createdAt
}

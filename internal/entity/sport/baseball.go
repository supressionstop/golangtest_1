package sport

import (
	"softpro6/internal/valueobject"
	"time"
)

type Baseball struct {
	rate      valueobject.Rate
	createdAt time.Time
}

func NewBaseball(rate valueobject.Rate, createdAt time.Time) *Baseball {
	return &Baseball{
		rate:      rate,
		createdAt: createdAt,
	}
}

func (s Baseball) Name() string {
	return "baseball"
}

func (s Baseball) Rate() valueobject.Rate {
	return s.rate
}

func (s Baseball) CreatedAt() time.Time {
	return s.createdAt
}

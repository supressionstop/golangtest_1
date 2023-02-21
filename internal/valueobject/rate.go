package valueobject

import (
	"github.com/shopspring/decimal"
)

type Rate struct {
	provider string
	value    decimal.Decimal
}

func NewRate(provider string, value float64) Rate {
	return Rate{
		provider: provider,
		value:    decimal.NewFromFloat(value),
	}
}

func NewRateFromString(rateValue, provider string) (Rate, error) {
	decimalValue, err := decimal.NewFromString(rateValue)
	if err != nil {
		return Rate{}, err
	}

	return Rate{
		provider: provider,
		value:    decimalValue,
	}, nil
}

func (r Rate) String() string {
	return r.value.String()
}

func (r Rate) Provider() string {
	return r.provider
}

func (r Rate) Value() decimal.Decimal {
	return r.value
}

func (r Rate) Equal(rate Rate) bool {
	return r.SameProvider(rate) && r.value.Equal(rate.Value())
}

func (r Rate) SameProvider(rate Rate) bool {
	return r.provider == rate.provider
}

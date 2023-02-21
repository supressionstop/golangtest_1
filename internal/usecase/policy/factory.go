package policy

import (
	"fmt"
	"softpro6/internal/entity/sport"
	"softpro6/internal/usecase"
	"strconv"
	"time"
)

type LineToSport func(line usecase.Line) (usecase.Sport, error)

func (f LineToSport) Export(line usecase.Line) (usecase.Sport, error) {
	rate, err := strconv.ParseFloat(line.Rate(), 64)
	if err != nil {
		return nil, err
	}
	t := time.Now()

	switch line.Sport() {
	case "baseball":
		return sport.NewBaseball(rate, t), nil
	default:
		return nil, fmt.Errorf("LineToSport - unknown sport: %q", line.Sport())
	}
}

package valueobject_test

import (
	"github.com/stretchr/testify/require"
	vo "softpro6/internal/valueobject"
	"testing"
)

func TestRate_String(t *testing.T) {
	tests := []struct {
		testVal string
	}{
		{testVal: "1"},
		{testVal: "1.01"},
		{testVal: "1.001"},
		{testVal: "1.0001"},
		{testVal: "1.00001"},
		{testVal: "1.000001"},
		{testVal: "1.0000001"},
		{testVal: "1.00000001"},
		{testVal: "1.000000001"},
		{testVal: "1.0000000001"},
		{testVal: "1.00000000001"},
		{testVal: "1.000000000001"},
		{testVal: "1.0000000000001"},
		{testVal: "1.00000000000001"},
		{testVal: "1.000000000000001"},
		{testVal: "13.000000000000001"},
		{testVal: "12.000000000000001"},
		{testVal: "199999999999.0000000000000001"},
		{testVal: "-199999999999.0000000000000001"},
		{testVal: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.testVal, func(t *testing.T) {
			r, err := vo.NewRateFromString(tt.testVal, tt.testVal)
			require.NoError(t, err)
			require.Equal(t, tt.testVal, r.String())
		})
	}
}

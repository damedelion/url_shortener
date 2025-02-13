package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowInt64(t *testing.T) {
	tests := []struct {
		x        int64
		y        int64
		expected int64
	}{
		{2, 0, 1},
		{2, 1, 2},
		{2, 3, 8},
		{3, 2, 9},
		{5, 4, 625},
		{10, 5, 100000},
		{0, 5, 0},
		{1, 1000, 1},
		{-2, 3, -8},
		{-3, 2, 9},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("x=%d, y=%d", tt.x, tt.y), func(t *testing.T) {
			result := PowInt64(tt.x, tt.y)
			assert.Equal(t, tt.expected, result)
		})
	}
}

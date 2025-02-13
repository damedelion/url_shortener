package base63

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBase63(t *testing.T) {
	tests := []struct {
		id        int64
		resLength int
		expected  string
		err       error
	}{
		{123456, 6, "000v6D", nil},
		{999, 4, "00fS", nil},
		{0, 4, "0000", nil},
		{62, 4, "000_", nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("id=%d, resLength=%d", tt.id, tt.resLength), func(t *testing.T) {
			result, err := ToBase63(tt.id, tt.resLength)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestBase63Table(t *testing.T) {
	tests := []struct {
		num      byte
		expected byte
		err      error
	}{
		{0, '0', nil},
		{9, '9', nil},
		{10, 'a', nil},
		{35, 'z', nil},
		{36, 'A', nil},
		{61, 'Z', nil},
		{62, '_', nil},
		{63, 0, fmt.Errorf("num is > 63")},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("num=%d", tt.num), func(t *testing.T) {
			result, err := base63Table(tt.num)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

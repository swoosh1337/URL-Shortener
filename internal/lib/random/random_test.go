package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name  string
		size  int
		count int // number of strings to generate and compare
	}{
		{name: "size = 1", size: 1, count: 100},
		{name: "size = 5", size: 5, count: 50},
		{name: "size = 10", size: 10, count: 30},
		{name: "size = 20", size: 20, count: 20},
		{name: "size = 30", size: 30, count: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strSet := make(map[string]bool, tt.count)
			for i := 0; i < tt.count; i++ {
				str := NewRandomString(tt.size)
				assert.Len(t, str, tt.size)
				strSet[str] = true
			}

			// Asserting on the count of unique strings generated
			assert.True(t, len(strSet) > 1, "Expected multiple unique strings, but got duplicates.")
		})
	}
}

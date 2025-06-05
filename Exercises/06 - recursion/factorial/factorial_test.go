package factorial_test

import (
	"github.com/stretchr/testify/assert"
	"roadmap-retos-programacion-mouredev/factorial"
	"testing"
)

func Test_Factorial(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "when n is 3, then the expected is 6",
			n:        3,
			expected: 6,
		},
		{
			name:     "when n is 5, then the expected is 120",
			n:        5,
			expected: 120,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := factorial.Factorial(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

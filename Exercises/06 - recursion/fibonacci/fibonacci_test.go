package fibonacci_test

import (
	"github.com/stretchr/testify/assert"
	"roadmap-retos-programacion-mouredev/fibonacci"
	"testing"
)

func Test_Fibonacci(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "when n is 0, then the result is 0",
			n:        0,
			expected: 0,
		},
		{
			name:     "when n is 1, then the result is 1",
			n:        1,
			expected: 1,
		},
		{
			name:     "when n is 5, then the result is 5",
			n:        5,
			expected: 5,
		},
		{
			name:     "when n is 10, then the result is 55",
			n:        10,
			expected: 55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fibonacci.Fibonacci(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

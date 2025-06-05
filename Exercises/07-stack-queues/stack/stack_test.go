package stack_test

import (
	"github.com/stretchr/testify/assert"
	"roadmap-retos-programacion-mouredev/stack"
	"testing"
)

func Test_StackPush(t *testing.T) {
	tests := []struct {
		name        string
		items       []string
		s           stack.Stack[string]
		expectedLen int
	}{
		{
			name:        "when any items are pushed, then the stack is empty",
			items:       nil,
			s:           stack.New[string](0),
			expectedLen: 0,
		},
		{
			name:        "when 1 item is pushed, then the size is 1",
			items:       []string{"item1"},
			s:           stack.New[string](1),
			expectedLen: 1,
		},
		{
			name:        "when multiple items are pushed, then the size is equal to the number of items",
			items:       []string{"item1", "item2"},
			s:           stack.New[string](2),
			expectedLen: 2,
		},
		{
			name:        "when the stack is full and a new item is pushed, then the item is not pushed",
			items:       []string{"item1", "item2"},
			s:           stack.New[string](1),
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, item := range tt.items {
				tt.s.Push(item)
			}
			assert.Equal(t, tt.expectedLen, tt.s.Size())
		})
	}
}

func Test_StackPop(t *testing.T) {
	tests := []struct {
		name string
		s    stack.Stack[string]
		test func(t *testing.T, s stack.Stack[string])
	}{
		{
			name: "when the stack is empty, then any items are popped",
			s:    stack.New[string](0),
			test: func(t *testing.T, s stack.Stack[string]) {
				_, ok := s.Pop()

				assert.False(t, ok)
				assert.Equal(t, 0, s.Size())
			},
		},
		{
			name: "when an item is popped, then the item is returned and the stack has 1 less item",
			s:    stack.New[string](2),
			test: func(t *testing.T, s stack.Stack[string]) {
				s.Push("item1")
				s.Push("item2")
				item, ok := s.Pop()

				assert.True(t, ok)
				assert.Equal(t, "item2", item)
				assert.Equal(t, 1, s.Size())
			},
		},
		{
			name: "when the stack is full and item is popped, then a new item can be pushed",
			s:    stack.New[string](2),
			test: func(t *testing.T, s stack.Stack[string]) {
				s.Push("item1")
				s.Push("item2")
				assert.Equal(t, 2, s.Size())

				_, ok := s.Pop()
				assert.True(t, ok)
				assert.Equal(t, 1, s.Size())

				ok = s.Push("item3")
				assert.True(t, ok)
				assert.Equal(t, 2, s.Size())
			},
		},
		{
			name: "when all items are popped, then the stack is empty",
			s:    stack.New[string](2),
			test: func(t *testing.T, s stack.Stack[string]) {
				s.Push("item1")
				s.Push("item2")
				assert.Equal(t, 2, s.Size())

				item, _ := s.Pop()
				assert.Equal(t, "item2", item)

				item, _ = s.Pop()
				assert.Equal(t, "item1", item)
				assert.Equal(t, 0, s.Size())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t, tt.s)
		})
	}
}

func Test_StackPeek(t *testing.T) {
	tests := []struct {
		name string
		s    stack.Stack[string]
		test func(t *testing.T, s stack.Stack[string])
	}{
		{
			name: "when the stack is empty, then any items are peeked",
			s:    stack.New[string](0),
			test: func(t *testing.T, s stack.Stack[string]) {
				_, ok := s.Peek()

				assert.False(t, ok)
				assert.Equal(t, 0, s.Size())
			},
		},
		{
			name: "when an item is peeked, then the item is returned and the stack has the same items",
			s:    stack.New[string](2),
			test: func(t *testing.T, s stack.Stack[string]) {
				s.Push("item1")
				s.Push("item2")
				item, ok := s.Peek()

				assert.True(t, ok)
				assert.Equal(t, "item2", item)
				assert.Equal(t, 2, s.Size())
			},
		},
		{
			name: "when the stack is full and item is peeked, then a new item cannot be pushed",
			s:    stack.New[string](2),
			test: func(t *testing.T, s stack.Stack[string]) {
				s.Push("item1")
				s.Push("item2")
				assert.Equal(t, 2, s.Size())

				_, ok := s.Peek()
				assert.True(t, ok)
				assert.Equal(t, 2, s.Size())

				ok = s.Push("item3")
				assert.False(t, ok)
			},
		},
		{
			name: "when peek is done multiple times against the same stack, then same item is returned all the times",
			s:    stack.New[string](2),
			test: func(t *testing.T, s stack.Stack[string]) {
				s.Push("item1")
				s.Push("item2")
				assert.Equal(t, 2, s.Size())

				item, _ := s.Peek()
				assert.Equal(t, "item2", item)

				item, _ = s.Peek()
				assert.Equal(t, "item2", item)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t, tt.s)
		})
	}
}

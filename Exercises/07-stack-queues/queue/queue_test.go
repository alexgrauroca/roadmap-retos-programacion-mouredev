package queue_test

import (
	"github.com/stretchr/testify/assert"
	"roadmap-retos-programacion-mouredev/queue"
	"testing"
)

func Test_QueueEnqueue_String(t *testing.T) {
	tests := []struct {
		name        string
		items       []string
		q           queue.Queue[string]
		expectedLen int
	}{
		{
			name:        "when any items enqueued, then the length is 0",
			items:       nil,
			q:           queue.New[string](0),
			expectedLen: 0,
		},
		{
			name:        "when one item enqueued, then the length is 1",
			items:       []string{"item1"},
			q:           queue.New[string](1),
			expectedLen: 1,
		},
		{
			name:        "when multiple items enqueued, then the length is equal to the number of items",
			items:       []string{"item1", "item2", "item3"},
			q:           queue.New[string](3),
			expectedLen: 3,
		},
		{
			name:        "when adding an extra item in a full queue, the last item is not added and the length remains the same",
			items:       []string{"item1", "item2", "item3", "item4", "item5"},
			q:           queue.New[string](4),
			expectedLen: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, item := range tt.items {
				tt.q.Enqueue(item)
			}
			assert.Equal(t, tt.expectedLen, tt.q.Size())
		})
	}
}

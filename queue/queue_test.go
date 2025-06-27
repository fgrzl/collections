package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue_EnqueueDequeue(t *testing.T) {
	// Arrange
	q := NewQueue[int]()

	// Act
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Assert
	require.Equal(t, 3, q.Length())

	v, ok := q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 1, v)

	v, ok = q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 2, v)

	v, ok = q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 3, v)

	require.True(t, q.IsEmpty())
}

func TestQueue_DequeueEmpty(t *testing.T) {
	// Arrange
	q := NewQueue[int]()

	// Act
	v, ok := q.Dequeue()

	// Assert
	require.False(t, ok)
	require.Equal(t, 0, v) // zero value for int
}

func TestQueue_Reset(t *testing.T) {
	// Arrange
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)

	// Act
	q.Reset()

	// Assert
	require.True(t, q.IsEmpty())
	require.Equal(t, 0, q.Length())

	// Reuse after reset
	q.Enqueue(42)
	v, ok := q.Dequeue()
	require.True(t, ok)
	require.Equal(t, 42, v)
}

func TestQueue_Reallocation(t *testing.T) {
	// Arrange
	q := NewQueue[int]()
	for i := 0; i < 4; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 2; i++ {
		_, _ = q.Dequeue()
	}

	// Act
	q.Enqueue(4)
	q.Enqueue(5)

	// Assert
	require.Equal(t, 4, q.Length())
	require.Equal(t, 0, q.Head()) // after reallocation, head is reset
	require.Equal(t, 4, q.Tail()) // tail is updated to new length
}

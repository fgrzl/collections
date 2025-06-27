package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack_PushPop(t *testing.T) {
	// Arrange
	s := NewStack[int]()

	// Act
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Assert
	require.Equal(t, 3, s.Length())

	v, ok := s.Pop()
	require.True(t, ok)
	require.Equal(t, 3, v)

	v, ok = s.Pop()
	require.True(t, ok)
	require.Equal(t, 2, v)

	v, ok = s.Pop()
	require.True(t, ok)
	require.Equal(t, 1, v)

	require.True(t, s.IsEmpty())
}

func TestStack_PopEmpty(t *testing.T) {
	// Arrange
	s := NewStack[int]()

	// Act
	v, ok := s.Pop()

	// Assert
	require.False(t, ok)
	require.Equal(t, 0, v) // zero value for int
}

func TestStack_Peek(t *testing.T) {
	// Arrange
	s := NewStack[int]()
	s.Push(10)

	// Act
	v, ok := s.Peek()

	// Assert
	require.True(t, ok)
	require.Equal(t, 10, v)
	require.Equal(t, 1, s.Length()) // ensure Peek didn't remove it
}

func TestStack_PeekEmpty(t *testing.T) {
	// Arrange
	s := NewStack[int]()

	// Act
	v, ok := s.Peek()

	// Assert
	require.False(t, ok)
	require.Equal(t, 0, v)
}

func TestStack_Reset(t *testing.T) {
	// Arrange
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)

	// Act
	s.Reset()

	// Assert
	require.True(t, s.IsEmpty())
	require.Equal(t, 0, s.Length())

	// Reuse after reset
	s.Push(99)
	v, ok := s.Pop()
	require.True(t, ok)
	require.Equal(t, 99, v)
}

func TestStack_WithCapacity(t *testing.T) {
	// Arrange
	s := NewStack[int](WithCapacity[int](10))

	// Act
	s.Push(1)

	// Assert
	require.Equal(t, 1, s.Length())
}

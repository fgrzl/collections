package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSet_Add(t *testing.T) {
	// Arrange
	set := NewHashSet[int]()

	// Act
	set.Add(1)

	// Assert
	assert.True(t, set.Contains(1), "Expected set to contain 1 after adding it")
}

func TestHashSet_Remove(t *testing.T) {
	// Arrange
	set := NewHashSet[int]()
	set.Add(1)

	// Act
	set.Remove(1)

	// Assert
	assert.False(t, set.Contains(1), "Expected set to not contain 1 after removing it")
}

func TestHashSet_Contains(t *testing.T) {
	// Arrange
	set := NewHashSet[int]()
	set.Add(1)

	// Act & Assert
	assert.True(t, set.Contains(1), "Expected set to contain 1")
	assert.False(t, set.Contains(2), "Expected set to not contain 2")
}

func TestHashSet_IsEmpty(t *testing.T) {
	// Arrange
	set := NewHashSet[int]()

	// Assert initial state
	assert.True(t, set.IsEmpty(), "Expected set to be empty initially")

	// Act
	set.Add(1)

	// Assert after adding
	assert.False(t, set.IsEmpty(), "Expected set to not be empty after adding an element")
}

func TestHashSet_Size(t *testing.T) {
	// Arrange
	set := NewHashSet[int]()

	// Assert initial size
	assert.Equal(t, 0, set.Size(), "Expected size to be 0 initially")

	// Act
	set.Add(1)
	set.Add(2)

	// Assert after adding
	assert.Equal(t, 2, set.Size(), "Expected size to be 2 after adding two elements")
}

func TestHashSet_Clear(t *testing.T) {
	// Arrange
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)

	// Act
	set.Clear()

	// Assert
	assert.True(t, set.IsEmpty(), "Expected set to be empty after clearing")
}

func TestHashSet_ToSlice(t *testing.T) {
	// Arrange
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)

	// Act
	slice := set.ToSlice()

	// Assert
	assert.Len(t, slice, 2, "Expected slice length to be 2")
	assert.Contains(t, slice, 1)
	assert.Contains(t, slice, 2)
}

func TestHashSet_ForEach(t *testing.T) {
	// Arrange
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	sum := 0

	// Act
	set.ForEach(func(item int) {
		sum += item
	})

	// Assert
	assert.Equal(t, 3, sum, "Expected sum to be 3")
}

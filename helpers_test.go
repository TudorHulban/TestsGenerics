package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertAtIndex(t *testing.T) {
	x := []int{1, 2, 4}
	insertAtIndex[int](&x, 2, 3)

	require.Equal(t, []int{1, 2, 3, 4}, x)
}

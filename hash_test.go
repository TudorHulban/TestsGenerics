package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssignments(t *testing.T) {
	root := newRoot(7)

	require.NoError(t, root.mapAssignments(), "assignments")
}

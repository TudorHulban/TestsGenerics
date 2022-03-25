package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssignments(t *testing.T) {
	root := newRoot(7)

	require.NoError(t, root.mapAssignments(), "assignments")
}

func TestAssignmentsDistribution(t *testing.T) {
	root := newRoot(7)
	n1 := newNode(9, root.id)

	root.registerNode(n1)

	a := hash.assignments(1, *root.getRing(root.getNodeData()))
	a.writeTo(os.Stdout)
}

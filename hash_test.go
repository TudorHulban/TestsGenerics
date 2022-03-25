package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssignments(t *testing.T) {
	root := newRoot(7)

	require.NoError(t, root.mapAssignments(), "assignments")
}

func TestAssignmentsDistribution(t *testing.T) {
	root := newRoot(7)

	chairs := 10
	maxReplicationFactor := 5

	output, errCr := os.Create("assignment_distribution")
	if errCr != nil {
		t.FailNow()
	}

	writeTo := output

	for i := 1; i <= chairs; i++ {
		if i == root.id {
			continue
		}

		root.registerNodeID(i)

		for f := 1; f < maxReplicationFactor; f++ {
			if f > len(*root.getRing(root.getNodeData())) {
				continue
			}

			a := hash.assignments(f, *root.getRing(root.getNodeData()))
			a.writeToWithFactor(writeTo, f)

			assert.NoError(t, hash.verifyFactor(f, a), fmt.Sprintf("factor verification for number nodes: %d", i))
		}
	}
}

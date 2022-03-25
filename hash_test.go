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

	chairs := 40
	maxReplicationFactor := 5

	output, errCr := os.Create("assignment_distribution")
	if errCr != nil {
		t.FailNow()
	}

	writeTo := output
	noPartitions := len(hash.partition())

	var totalCases int
	var noFailures int

	for i := 1; i < chairs; i++ {
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

			if !assert.NoError(t, hash.verifyFactor(f, noPartitions, a), fmt.Sprintf("factor verification for number nodes: %d", i)) {
				noFailures++
			}

			totalCases++
		}
	}

	t.Logf("\nFailed cases: %d from a total of %d.\n", noFailures, totalCases)
}

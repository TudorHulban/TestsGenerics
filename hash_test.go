package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssignments(t *testing.T) {
	root := newRoot(7)

	require.NoError(t, root.mapAssignments(), "assignments")
}

func TestAssignmentsDistribution(t *testing.T) {
	root := newRoot(7)

	chairs := 2
	// maxReplicationFactor := 1

	// output, errCr := os.Create("assignment_distribution")
	// if errCr != nil {
	// 	t.FailNow()
	// }

	// writeTo := output
	// noPartitions := len(hash.partition())

	// var totalCases int
	// var noFailures int

	for i := 1; i < chairs; i++ {
		if i == root.id {
			continue
		}

		fmt.Printf("chairs: %d", i)

		root.registerNodeID(i)

		// a := hash.assign(maxReplicationFactor, *root.getRing(root.getNodeData()))
		// a.writeToWithFactor(writeTo, maxReplicationFactor)

		// for f := 1; f <= maxReplicationFactor; f++ {

		// 	if f > len(*root.getRing(root.getNodeData())) {
		// 		continue
		// 	}

		// 	a := hash.assign(f, *root.getRing(root.getNodeData()))
		// 	a.writeToWithFactor(writeTo, f)

		// 	if !assert.NoError(t, hash.verifyFactor(f, noPartitions, a), fmt.Sprintf("factor verification for number nodes: %d", i)) {
		// 		noFailures++
		// 	}

		// 	totalCases++
		// }
	}

	// t.Logf("\nFailed cases: %d from a total of %d.\n", noFailures, totalCases)
}

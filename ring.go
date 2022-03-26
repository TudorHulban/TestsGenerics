package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type NodeData struct {
	Ranges Partitions
	ID     int
}


type ring struct{
	// nodes []*NodeData,
	partitions Partitions,
}

func newRing(nodes ...*NodeData) *ring {
	if len(nodes) == 0 {
		return nil
	}

	var res ring

	res = append(res, nodes...)

	return &res
}

func (r ring) getAssignments() Assignments {
	var partitions []string
	p := NewPartitions(1) // TODO: remove factor

	for _, partition := range p {
		for i := 0; i < partition.Factor; i++ {
			partitions = append(partitions, partition.Range)
		}
	}

	fmt.Printf("partitions: %d, factor: %d\n", len(partitions), 1)

loop:
	for i := 0; i < len(partitions); i++ {
		for _, node := range ring {
			if !slices.Contains(node.Partitions, partitions[i]) {
				node.Partitions = append(node.Partitions, partitions[i])
			}

			if i == len(partitions)-1 {
				break loop
			}

			i++
		}
	}

	res := make(map[int][]string)

	for nodeID, partitions := range ring {
		res[nodeID] = partitions.Partitions
	}

	return res
}

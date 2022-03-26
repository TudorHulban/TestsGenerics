package main

import (
	"fmt"
	"io"
	"strings"
)

type assignment struct {
	nodeID int
	ranges Partitions
}

type Assignments map[int]Partitions // key is node ID

func (a Assignments) writeTo(w io.Writer) {
	var res []string

	res = append(res, fmt.Sprintf("Ranges: %s.", NewPartitions(1).getRanges()))

	for nodeID, partitions := range a {
		res = append(res, fmt.Sprintf("Node ID %d: %v", nodeID, partitions.getRanges()))
	}

	res = append(res, "\n")

	w.Write([]byte(strings.Join(res, "\n")))
}

func (a Assignments) writeToWithFactor(w io.Writer, factor int) {
	w.Write([]byte(fmt.Sprintf("Replication factor is %d.\n", factor)))

	a.writeTo(w)
}

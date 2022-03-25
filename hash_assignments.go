package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type assignment struct {
	nodeID     int
	partitions []string
}

type Assignments map[int][]string

func (a Assignments) writeTo(w io.Writer) {
	var res []string

	res = append(res, fmt.Sprintf("Partitions: %s.", hash.partition()))

	assign := a.sorted()

	for _, assignment := range assign {
		res = append(res, fmt.Sprintf("Node ID %d: %v", assignment.nodeID, assignment.partitions))
	}

	res = append(res, "\n")

	w.Write([]byte(strings.Join(res, "\n")))
}

func (a Assignments) writeToWithFactor(w io.Writer, factor int) {
	w.Write([]byte(fmt.Sprintf("Replication factor is %d.\n", factor)))
	a.writeTo(w)
}

func (a Assignments) sorted() []assignment {
	var res []assignment

	for nodeID, partitions := range a {
		res = append(res, assignment{
			nodeID:     nodeID,
			partitions: partitions,
		})
	}

	asc := func(i, j int) bool {
		return res[i].nodeID < res[j].nodeID
	}

	sort.Slice(res, asc)

	return res
}

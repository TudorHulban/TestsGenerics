package main

import (
	"fmt"
	"strconv"
)

type hasher func(string) string

func (h hasher) hashNode(n node) string {
	return h(strconv.Itoa(n.id))
}

func (h hasher) verifyFactor(factor, noPartitions int, a Assignments) error {
	var assignments []string

	for _, partitions := range a {
		assignments = append(assignments, partitions...)
	}

	occ := occurences[string](assignments)

	if noPartitions > len(occ) {
		return fmt.Errorf("not all partitions were mapped. missing %d", noPartitions-len(occ))
	}

	for partition, f := range occ {
		if f < factor {
			return fmt.Errorf("for partition '%s' factor is only %d versus required %d", partition, f, factor)
		}
	}

	return nil
}

var hash hasher = func(s string) string {
	return (s + "000000000")[:10]
}

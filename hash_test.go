package main

import (
	"fmt"
	"testing"
)

func TestAssignments(t *testing.T) {
	hash.assignments(2, 5)
}

func TestChop(t *testing.T) {
	partitions := hash.chop(5, 3)

	fmt.Println(partitions)
}

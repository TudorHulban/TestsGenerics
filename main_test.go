package main

import (
	"fmt"
	"testing"
)

func TestHashing(t *testing.T) {
	n := node{
		id: 34,
	}

	fmt.Printf(hash.hashNode(n) + "\n")
}

package main

import (
	"fmt"
	"io"
	"strings"
)

type Assignments map[int][]string // key is node identification

func (a Assignments) writeTo(w io.Writer) {
	var res []string

	for nodeID, assignment := range a {
		res = append(res, fmt.Sprintf("Node ID %d: %s", nodeID, assignment))
	}

	res = append(res, "\n")

	w.Write([]byte(strings.Join(res, "\n")))
}

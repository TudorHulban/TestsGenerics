package main

import (
	"fmt"
	"io"
	"strings"
)

type neighbors struct {
	next     []*NodeData // ordered, ascending
	previous []*NodeData // ordered, ascending
}

func newNeighbors() *neighbors {
	return &neighbors{
		next:     make([]*NodeData, 0),
		previous: make([]*NodeData, 0),
	}
}

func (n *neighbors) appendToPrevious(no *NodeData) {
	if len(n.previous) == 0 {
		n.previous = append(n.previous, no)

		return
	}

	for i := len(n.previous) - 1; i >= 0; i-- {
		if no.ID > n.previous[i].ID {
			insertAtIndex[*NodeData](&n.previous, i, no)

			return
		}
	}

	insertAtIndex[*NodeData](&n.previous, 0, no)
}

func (n *neighbors) appendToNext(no *NodeData) {
	if len(n.next) == 0 {
		n.next = append(n.next, no)

		return
	}

	for i := 0; i < len(n.next); i++ {
		if no.ID < n.next[i].ID {
			insertAtIndex[*NodeData](&n.next, i, no)

			return
		}
	}

	n.next = append(n.next, no)
}

func (n neighbors) neighborsTo(w io.Writer) {
	var res []string

	info := func(ix int, no *NodeData) {
		res = append(res, fmt.Sprintf("Element: %d with ID: %d.", ix, no.ID))
	}

	res = append(res, fmt.Sprintf("Previous Set(%d):", len(n.previous)))
	forEach[*NodeData](n.previous, info)

	res = append(res, fmt.Sprintf("Next Set(%d):", len(n.next)))
	forEach[*NodeData](n.next, info)

	res = append(res, "\n")

	w.Write([]byte(strings.Join(res, "\n")))
}

func (n *neighbors) getRing(localNode *NodeData) *ring {
	var res ring

	res = append(res, n.previous...)
	res = append(res, localNode)
	res = append(res, n.next...)

	return &res
}

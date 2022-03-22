package main

import (
	"fmt"
	"io"
	"strings"
)

type neighbors struct {
	next     []*nodeData // ordered, ascending
	previous []*nodeData // ordered, ascending
}

func newNeighbors() *neighbors {
	return &neighbors{
		next:     make([]*nodeData, 0),
		previous: make([]*nodeData, 0),
	}
}

func (n *neighbors) appendToPrevious(no *nodeData) {
	if len(n.previous) == 0 {
		n.previous = append(n.previous, no)

		return
	}

	for i := len(n.previous) - 1; i >= 0; i-- {
		if no.id > n.previous[i].id {
			insertAtIndex[*nodeData](&n.previous, i, no)
			return
		}
	}
}

func (n *neighbors) appendToNext(no *nodeData) {
	if len(n.next) == 0 {
		n.next = append(n.next, no)

		return
	}

	for i := 0; i < len(n.next); i++ {
		if no.id < n.next[i].id {
			insertAtIndex[*nodeData](&n.next, i, no)
			return
		}
	}
}

func (n neighbors) neighborsTo(w io.Writer) {
	var res []string

	info := func(ix int, no *nodeData) {
		res = append(res, fmt.Sprintf("Element: %d with ID: %d.\n", ix, no.id))
	}

	res = append(res, fmt.Sprintf("Previous Set(%d):", len(n.previous)))
	forEach[*nodeData](n.previous, info)

	res = append(res, fmt.Sprintf("Next Set(%d):", len(n.next)))
	forEach[*nodeData](n.next, info)

	w.Write([]byte(strings.Join(res, "\n")))
}

func (n *neighbors) getRing() *ring {
	var res ring

	res = append(res, n.previous...)
	res = append(res, n.next...)

	return &res
}

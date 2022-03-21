package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type node struct {
	next     []*node // ordered, ascending
	previous []*node // ordered, ascending

	cache      *cache
	partitions []string

	id     int // node IDs are known
	factor int // redundancy factor
	rootID int
}

func newNode(id, rootID int) *node {
	return &node{
		id:     id,
		rootID: rootID,
		factor: 2,

		next:     make([]*node, 0),
		previous: make([]*node, 0),
	}
}

func newRoot(id int) *node {
	n := newNode(id, id)

	n.partitions = hash.partition()

	return n
}

func (n node) isRoot() bool {
	return n.id == n.rootID
}

func (n node) getID() int {
	return n.id
}

func (n *node) registerNode(no *node) error {
	if n.id > no.id {
		n.appendToPrevious(no)
		return nil
	}

	if no.id == n.id {
		return fmt.Errorf("node to register has the same ID(%d) with curent node", no.id)
	}

	n.appendToNext(no)

	return nil
}

func (n *node) appendToPrevious(no *node) {
	if len(n.previous) == 0 {
		n.previous = append(n.previous, no)

		return
	}

	for i := len(n.previous) - 1; i >= 0; i-- {
		if no.id > n.previous[i].id {
			insertAtIndex[*node](&n.previous, i, no)
			return
		}
	}
}

func (n *node) appendToNext(no *node) {
	if len(n.next) == 0 {
		n.next = append(n.next, no)

		return
	}

	for i := 0; i < len(n.next); i++ {
		if no.id < n.next[i].id {
			insertAtIndex[*node](&n.next, i, no)
			return
		}
	}
}

func (n node) neighborsTo(w io.Writer) {
	var res []string

	res = append(res, fmt.Sprintf("Node ID; %d", n.id))
	res = append(res, fmt.Sprintf("Previous Set(%d):", len(n.previous)))

	info := func(ix int, no *node) {
		res = append(res, fmt.Sprintf("Element: %d with ID: %d.\n", ix, no.id))
	}

	forEach[*node](n.previous, info)

	res = append(res, fmt.Sprintf("Next Set(%d):", len(n.next)))

	forEach[*node](n.next, info)

	w.Write([]byte(strings.Join(res, "\n")))
}

func (n node) listenOn() string {
	id := "000" + strconv.Itoa(n.id)

	return "8" + id[len(id)-3:]
}

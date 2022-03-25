package main

import (
	"fmt"
	"io"
	"strings"
)

type neighbors struct {
	previous []*NodeData // ordered, ascending
	local    *NodeData
	next     []*NodeData // ordered, ascending
}

func newNeighbors(localID int) *neighbors {
	return &neighbors{
		previous: make([]*NodeData, 0),

		// duplicated info
		local: &NodeData{
			ID:         localID,
			Partitions: hash.partition(),
		},

		next: make([]*NodeData, 0),
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

func (n *neighbors) getPartitionsFor(id int, w io.Writer) {
	partitions := n.getNeighborData(id).Partitions

	fmt.Printf("get partitions: %s", partitions)

	w.Write([]byte(strings.Join(partitions, ",")))
}

func (n *neighbors) getNeighborData(id int) *NodeData {
	getNode := func(s []*NodeData) *NodeData {
		for _, node := range s {
			if node.ID == id {
				return node
			}
		}

		return nil
	}

	if id < n.local.ID {
		return getNode(n.previous)
	}

	if id == n.local.ID {
		return n.local
	}

	return getNode(n.next)
}

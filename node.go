package main

type node struct {
	next     []*node // ordered, ascending
	previous []*node // ordered, ascending

	id     string
	socket string
}

func (n *node) registerNode(no *node) {
	if no.id < low[*node](n.next).id {

	}
}

func (n *node) appendToPrevious(no *node) {
	for i := len(n.previous) - 1; i >= 0; i-- {
		if no.id > n.previous[i].id {
			insertAtIndex[*node](&n.previous, i, no)
		}
	}
}

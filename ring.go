package main

type NodeData struct {
	Partitions []string
	ID         int
}

type ring []*NodeData // would be sent in announciation response

func newRing(nodes ...*NodeData) *ring {
	if len(nodes) == 0 {
		return nil
	}

	var res ring

	res = append(res, nodes...)

	return &res
}

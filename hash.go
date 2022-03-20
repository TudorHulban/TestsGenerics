package main

import "strconv"

type hasher func(string) string

func (h hasher) hashNode(n node) string {
	return h(strconv.Itoa(n.id))
}

func (h hasher) partition() []string {
	return []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
}

var hash hasher = func(s string) string {
	return (s + "000000000")[:10]
}

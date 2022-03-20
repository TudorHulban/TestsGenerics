package main

import (
	"fmt"
	"math"
	"strconv"
)

type hasher func(string) string

func (h hasher) hashNode(n node) string {
	return h(strconv.Itoa(n.id))
}

func (h hasher) partition() []string {
	return []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
}

func (h hasher) assignments(factor, nodes int) map[int][]string {
	res := make(map[int][]string)

	total := len(h.partition())
	each := factor * int(math.Floor(float64(total)/float64(nodes)))

	if each > total {
		each = total
	}

	fmt.Println(each)

	return res
}

func (h hasher) chop(each, nodes int) map[int][]string {
	res := make(map[int][]string)

	var positionPartition int
	lengthPartition := len(h.partition())

	for i := 0; i < nodes; i++ {
		var buf []string

		for j := 0; j < each; j++ {
			buf = append(buf, h.partition()[positionPartition])

			positionPartition++
			if positionPartition > lengthPartition-1 {
				positionPartition = 0
			}
		}

		res[i] = buf
	}

	return res
}

var hash hasher = func(s string) string {
	return (s + "000000000")[:10]
}
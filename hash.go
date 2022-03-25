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

func (h hasher) assignments(factor int, r ring) Assignments {
	lengthPartition := len(h.partition())

	if factor == len(r) {
		return h.chop(lengthPartition, factor, r)
	}

	each := factor * int(1+math.Floor(float64(lengthPartition)/float64(len(r))))

	extra := func() int {
		if lengthPartition%(each+1) != 0 {
			return 1
		}

		return 0
	}

	each = each + extra()

	if each > lengthPartition {
		each = lengthPartition
	}

	if each == 0 {
		each = 1
	}

	return h.chop(each, factor, r)
}

func (h hasher) chop(each, factor int, r ring) Assignments {
	res := make(map[int][]string)

	var positionPartition int
	lengthPartition := len(h.partition())

	for ix, nodeData := range r {
		var buf []string

		for j := 0; j < each; j++ {
			buf = append(buf, h.partition()[positionPartition])

			positionPartition++
			if positionPartition > lengthPartition-1 {
				positionPartition = 0
			}
		}

		if ix == len(r)-1 && factor > 1 && lengthPartition%each != 0 {
			buf = append(buf, h.partition()[positionPartition:]...)
		}

		res[nodeData.ID] = removeDups[string](buf)
	}

	return res
}

func (h hasher) verifyFactor(factor, noPartitions int, a Assignments) error {
	var assign []string

	for _, partitions := range a {
		assign = append(assign, partitions...)
	}

	occ := occurences[string](assign)

	if noPartitions > len(occ) {
		return fmt.Errorf("not all partitions were mapped. missing %d", noPartitions-len(occ))
	}

	for partition, f := range occ {
		if f < factor {
			return fmt.Errorf("for partition '%s' factor is only %d versus required %d", partition, f, factor)
		}
	}

	return nil
}

var hash hasher = func(s string) string {
	return (s + "000000000")[:10]
}

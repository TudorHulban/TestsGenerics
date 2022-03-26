package main

import "golang.org/x/exp/slices"

type Partition struct {
	Range  string
	Factor int
}

type Partitions []Partition

func NewPartitions(f int) Partitions {
	return []Partition{
		Partition{"0", f},
		Partition{"1", f},
		Partition{"2", f},
		Partition{"3", f},
		Partition{"4", f},
		Partition{"5", f},
		Partition{"6", f},
		Partition{"7", f},
		Partition{"8", f},
		Partition{"9", f},
	}
}

func (p Partitions) getRanges() []string {
	var res []string

	for _, partition := range p {
		res = append(res, partition.Range)
	}

	slices.Sort[string](res)

	return res
}

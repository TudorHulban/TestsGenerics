package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/TudorHulban/log"
)

type node struct {
	*neighbors

	cache      *cache
	partitions []string

	l *log.Logger

	id     int // node IDs are known
	rootID int
	factor int // redundancy factor
}

func newNode(id, rootID int) *node {
	return &node{
		neighbors: newNeighbors(id),
		id:        id,
		rootID:    rootID,
		factor:    2,
		l:         log.NewLogger(log.DEBUG, os.Stdout, true),
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

func (n *node) getNodeData() *NodeData {
	return &NodeData{
		ID:         n.id,
		Partitions: n.partitions,
	}
}

func (n *node) getNeighboorData(id int) *NodeData {
	getNode := func(s []*NodeData) *NodeData {
		for _, node := range s {
			if node.ID == id {
				return node
			}
		}

		return nil
	}

	if id < n.id {
		return getNode(n.previous)
	}

	if id == n.id {
		return n.getNodeData()
	}

	return getNode(n.next)
}

func (n *node) registerNode(no *node) error {
	if n.getNeighboorData(no.id) != nil {
		return fmt.Errorf("node with ID: %d is already registered", no.id)
	}

	if n.id > no.id {
		n.appendToPrevious(no.getNodeData())

		return nil
	}

	if no.id == n.id {
		return fmt.Errorf("node to register has the same ID(%d) with curent node", no.id)
	}

	n.appendToNext(no.getNodeData())

	return nil
}

func (n *node) registerNodeID(id int) error {
	no := node{
		id: id,
	}

	errReg := n.registerNode(&no)
	if errReg != nil {
		return errReg
	}

	return n.mapAssignments()
}

func (n node) listenFor(id int) string {
	sock := "000" + strconv.Itoa(id)

	return "127.0.0.1:8" + sock[len(sock)-3:]
}

func (n node) listenOn() string {
	return n.listenFor(n.id)
}

func (n node) announceTo(id int) error {
	url := "http://" + n.listenFor(id) + "/" + urlAnnounce + "/" + strconv.Itoa(n.id)

	fmt.Printf("Announcing node to URL: %s", url)

	req, errReq := http.NewRequest("GET", url, nil)
	if errReq != nil {
		return fmt.Errorf("announceTo: %w", errReq)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	var client http.Client

	resp, errCall := client.Do(req)
	if errCall != nil {
		return fmt.Errorf("announce on URL: %s gives: %w", url, errCall)
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return fmt.Errorf("status on announce to URL: %s is: %s", url, resp.Status)
	}

	return nil
}

func (n *node) mapAssignments() error {
	ring := n.getRing(n.getNodeData())
	assignments := hash.assignments(2, *ring)

	if len(*ring) != len(assignments) {
		return fmt.Errorf("different number of nodes (%d) versus assignments (%d)", len(*ring), len(assignments))
	}

	for _, nodeData := range *ring {
		data, exists := assignments[nodeData.ID]
		if !exists {
			return fmt.Errorf("assignment for node with ID: %d not found", nodeData.ID)
		}

		if len(data) == 0 {
			return fmt.Errorf("no partitions for node with ID: %d", nodeData.ID)
		}

		n.getNeighboorData(nodeData.ID).Partitions = []string{}
		n.getNeighboorData(nodeData.ID).Partitions = append(n.getNeighboorData(nodeData.ID).Partitions, data...)
	}

	return nil
}

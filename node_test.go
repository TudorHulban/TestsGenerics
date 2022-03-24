package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppend(t *testing.T) {
	root := newRoot(7)

	n1 := newNode(5, root.getID())

	root.appendToPrevious(n1.getNodeData())
	require.Equal(t, 1, len(root.previous), "previous")

	n2 := newNode(9, root.getID())

	root.appendToNext(n2.getNodeData())
	require.Equal(t, 1, len(root.next), "next")
}

func TestListenSock(t *testing.T) {
	root := newRoot(7)

	require.Equal(t, "127.0.0.1:8007", root.listenOn())
}

func TestRegister(t *testing.T) {
	root := newRoot(7)

	n1 := newNode(5, root.getID())

	require.NoError(t, root.registerNode(n1))
	require.Equal(t, 1, len(root.previous), "previous")
	require.Equal(t, 2, len(*root.getRing(root.getNodeData())))

	n2 := newNode(9, root.getID())
	require.NoError(t, root.registerNode(n2))

	n3 := newNode(8, root.getID())
	require.NoError(t, root.registerNode(n3))

	require.Equal(t, 2, len(root.next), "next")
	require.Equal(t, 4, len(*root.getRing(root.getNodeData())), "ring")

	root.neighborsTo(os.Stdout)

	require.NoError(t, root.mapAssignments(), "assignments")
	require.Greater(t, len(root.next[0].Partitions), 0, "partitions assignment")
}

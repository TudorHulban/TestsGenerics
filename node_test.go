package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppend(t *testing.T) {
	root := newRoot(7)

	n1 := newNode(5)

	root.appendToPrevious(n1)
	require.Equal(t, 1, len(root.previous), "previous")

	n2 := newNode(9)

	root.appendToNext(n2)
	require.Equal(t, 1, len(root.next), "next")
}

func TestListenSock(t *testing.T) {
	root := newRoot(7)

	require.Equal(t, "8007", root.listenOn())
}

func TestRegister(t *testing.T) {
	root := newRoot(7)

	n1 := newNode(5)

	root.registerNode(n1)
	require.Equal(t, 1, len(root.previous), "previous")

	n2 := newNode(9)
	root.registerNode(n2)

	n3 := newNode(8)
	root.registerNode(n3)

	require.Equal(t, 2, len(root.next), "next")

	root.neighborsTo(os.Stdout)
}

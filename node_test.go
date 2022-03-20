package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppend(t *testing.T) {
	root := newNode(7, true)

	n1 := newNode(5, false)

	root.appendToPrevious(n1)
	require.Equal(t, 1, len(root.previous), "previous")

	n2 := newNode(9, false)

	root.appendToNext(n2)
	require.Equal(t, 1, len(root.next), "next")
}

func TestListenSock(t *testing.T) {
	root := newNode(7, true)

	require.Equal(t, "8007", root.listenOn())
}

func TestRegister(t *testing.T) {
	root := newNode(7, true)

	n1 := newNode(5, false)

	root.registerNode(n1)
	require.Equal(t, 1, len(root.previous), "previous")

	n2 := newNode(9, false)

	root.registerNode(n2)
	require.Equal(t, 1, len(root.next), "next")

	root.neighborsTo(os.Stdout)
}

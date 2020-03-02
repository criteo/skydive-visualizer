package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopEmpty(t *testing.T) {
	g := Top(Graph{}, 1, 0)
	require.Equal(t, Graph{}, g)
}

func TestTopSimple(t *testing.T) {
	g := makeG(t, `
		{"0":0}
		{"0":1}
		{"0":2}
		{"0":3}
		0 -> 1 {"0":10, "1":11}
		2 -> 3 {"0":15, "1":5}
	`)
	g = Top(g, 1, 0)

	require.Equal(t, makeG(t, `
		{"0":2}
		{"0":3}
		0 -> 1 {"0":15, "1":5}
	`), g)
}

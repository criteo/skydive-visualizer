package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupSimple(t *testing.T) {
	g := makeG(t, `
		{"1":1, "2":1}
		{"1":1, "2":2}
		{"1":3, "2":3}
		0 -> 2 {"1":5}
		1 -> 2 {"1":10}
	`)

	g2 := Group(g, []Attribute{1}, []Attribute{2})

	require.Equal(t, makeG(t, `
		{"1":1}
		{"2":3}
		0 -> 1 {"1": 15}
	`), g2)
}

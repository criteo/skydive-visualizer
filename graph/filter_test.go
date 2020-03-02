package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	g := makeG(t, `
		{}
		{}
		{}
		{}
		0 -> 2 {}
		1 -> 2 {}
		1 -> 3 {}
	`)
	g.SetNodeAttr(0, 1, "ok")
	g.SetNodeAttr(1, 1, "not_ok")
	g.SetNodeAttr(2, 1, "ok2")
	g.SetNodeAttr(3, 1, "not_ok2")

	g2 := Filter(g,
		map[Attribute]string{1: "ok"},
		map[Attribute]string{1: "ok2"},
	)

	expected := g.New()
	expected.CreateNode(map[Attribute]string{
		1: "ok",
	})
	expected.CreateNode(map[Attribute]string{
		1: "ok2",
	})
	expected.AddEdge(Edge{
		FromNode: 0,
		ToNode:   1,
		Attrs:    EdgeAttributes{},
	})

	require.Equal(t, expected, g2)
}

func TestFilterAttrDoesNotExistsEmpty(t *testing.T) {
	g := makeG(t, `
		{}
		{}
		0 -> 1 {}
	`)

	g2 := Filter(g,
		map[Attribute]string{1: ""},
		map[Attribute]string{42: ""},
	)

	require.Equal(t, g, g2)
}

func TestFilterAttrDoesNotExistsNotEmpty(t *testing.T) {
	g := makeG(t, `
		{}
		{}
		0 -> 1 {}
	`)

	g2 := Filter(g,
		map[Attribute]string{1: "ok"},
		map[Attribute]string{42: "ok"},
	)

	require.Equal(t, g.New(), g2)
}

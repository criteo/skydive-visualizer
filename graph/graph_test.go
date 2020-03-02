package graph

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func makeG(t *testing.T, str string) Graph {
	g := Graph{}
	lines := strings.Split(str, "\n")

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		if !strings.Contains(l, "->") {
			// node

			attrs := map[Attribute]int{}
			err := json.Unmarshal([]byte(strings.TrimSpace(l)), &attrs)
			if err != nil {
				t.Fatalf("bad graph line: %s", err)
			}
			g.AddNode(Node{Attrs: attrs})
		} else {
			// edge

			var left, right, attrs string
			p1 := strings.Split(l, "->")
			if len(p1) != 2 {
				t.Fatalf("bad graph line: %q", l)
			}

			left = strings.TrimSpace(p1[0])

			p2 := strings.SplitN(strings.TrimSpace(p1[1]), " ", 2)
			if len(p2) != 2 {
				t.Fatalf("bad graph line: %q", l)
			}

			right = strings.TrimSpace(p2[0])
			attrs = strings.TrimSpace(p2[1])

			fromNode, err := strconv.Atoi(left)
			if err != nil {
				t.Fatalf("bad graph line: %s", err)
			}
			toNode, err := strconv.Atoi(right)
			if err != nil {
				t.Fatalf("bad graph line: %s", err)
			}

			linkAttrs := map[Attribute]float64{}
			err = json.Unmarshal([]byte(attrs), &linkAttrs)
			if err != nil {
				t.Fatalf("bad graph line: %s", err)
			}
			g.AddEdge(Edge{
				FromNode: fromNode,
				ToNode:   toNode,
				Attrs:    linkAttrs,
			})
		}
	}

	return g
}

func TestGraphFromString(t *testing.T) {
	g := makeG(t, `
		{"0":1, "1":2}
		{"0":1, "1":3}
		{"0":1, "1":4}
		{"0":1, "1":5}
		0 -> 1 {"0":10, "1":11}
		2 -> 3 {"0":15, "1":17}
	`)

	require.Equal(t, Graph{
		Nodes: []Node{
			{Attrs: map[Attribute]int{0: 1, 1: 2}},
			{Attrs: map[Attribute]int{0: 1, 1: 3}},
			{Attrs: map[Attribute]int{0: 1, 1: 4}},
			{Attrs: map[Attribute]int{0: 1, 1: 5}},
		},
		Edges: []Edge{
			{FromNode: 0, ToNode: 1, Attrs: map[Attribute]float64{0: 10, 1: 11}},
			{FromNode: 2, ToNode: 3, Attrs: map[Attribute]float64{0: 15, 1: 17}},
		},
	}, g)
}

func TestGraphAttributes(t *testing.T) {
	g := Graph{}

	g.CreateNode(map[Attribute]string{})
	g.CreateNode(map[Attribute]string{})
	g.CreateNode(map[Attribute]string{})

	g.SetNodeAttr(0, 1, "ok")
	g.SetNodeAttr(1, 1, "ok")
	g.SetNodeAttr(2, 3, "ok2")

	require.Equal(t, "ok", g.NodeAttr(0, 1))
	require.Equal(t, "ok", g.NodeAttr(1, 1))
	require.Equal(t, "ok2", g.NodeAttr(2, 3))
}

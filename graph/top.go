package graph

import "sort"

func Top(g Graph, n int, dim Attribute) Graph {
	if len(g.Edges) <= n {
		return g
	}

	out := g.New()
	out.Edges = make([]Edge, len(g.Edges))
	copy(out.Edges, g.Edges)

	sort.Slice(out.Edges, func(i, j int) bool {
		a := out.Edges[i].Attrs[dim]
		b := out.Edges[j].Attrs[dim]

		return a > b
	})

	out.Edges = out.Edges[:n]

	for i, link := range out.Edges {
		fn := out.AddNode(g.Nodes[link.FromNode])
		tn := out.AddNode(g.Nodes[link.ToNode])

		out.Edges[i].FromNode = fn
		out.Edges[i].ToNode = tn
	}

	return out
}

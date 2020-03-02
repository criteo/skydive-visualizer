package graph

import (
	"sort"
)

func Group(g Graph, fromDims, toDims []Attribute) Graph {
	out := g.New()

	edges := make([]Edge, 0, len(g.Edges))
	for _, e := range g.Edges {
		edges = append(edges, e)
	}

	sort.Slice(edges, func(i, j int) bool {
		a := edges[i]
		b := edges[j]

		an := g.Nodes[a.FromNode]
		bn := g.Nodes[b.FromNode]

		for _, d := range fromDims {
			av := an.Attrs[d]
			bv := bn.Attrs[d]

			if av == bv {
				continue
			}

			return av < bv
		}

		return false
	})

	var gedges []Edge

	var lastValues NodeAttributes
	var lastEdgeAttrs EdgeAttributes
	nodeIdx := 0
	for _, edge := range edges {
		fn := g.Nodes[edge.FromNode]

		fv := dimValues(fn.Attrs, fromDims)

		if fv.Equals(lastValues) && lastEdgeAttrs != nil {
			for k, v := range edge.Attrs {
				lastEdgeAttrs[k] += v
			}
			continue
		}

		nodeIdx++
		fgn := out.AddNode(Node{
			Attrs: fv,
		})
		lastValues = fv

		linkAttrs := edge.Attrs.Copy()

		gedges = append(gedges, Edge{
			FromNode: fgn,
			ToNode:   edge.ToNode,
			Attrs:    linkAttrs,
		})
		lastEdgeAttrs = linkAttrs
	}

	sort.Slice(gedges, func(i, j int) bool {
		a := edges[i]
		b := edges[j]

		an := g.Nodes[a.ToNode]
		bn := g.Nodes[b.ToNode]

		for _, d := range toDims {
			av := an.Attrs[d]
			bv := bn.Attrs[d]

			if av == bv {
				continue
			}

			return av < bv
		}

		return false
	})

	var lastFNode, lastTNode int
	lastValues = nil
	lastEdgeAttrs = nil

	for _, edge := range gedges {
		tn := g.Nodes[edge.ToNode]

		tv := dimValues(tn.Attrs, toDims)
		dimEq := lastValues.Equals(tv) && len(toDims) > 0

		if dimEq && edge.FromNode == lastFNode && lastEdgeAttrs != nil {
			for k, v := range edge.Attrs {
				lastEdgeAttrs[k] += v
			}
			continue
		}

		var tgn int
		if dimEq {
			tgn = lastTNode
		} else {
			nodeIdx++
			tgn = out.AddNode(Node{
				Attrs: tv,
			})
			lastValues = tv
			lastTNode = tgn
		}

		linkAttrs := edge.Attrs.Copy()

		out.AddEdge(Edge{
			FromNode: edge.FromNode,
			ToNode:   tgn,
			Attrs:    linkAttrs,
		})
		lastEdgeAttrs = linkAttrs
		lastFNode = edge.FromNode
	}

	return out
}

func dimValues(a NodeAttributes, dims []Attribute) NodeAttributes {
	res := NodeAttributes{}
	for _, d := range dims {
		res[d] = a[d]
	}
	return res
}

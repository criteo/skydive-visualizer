package graph

func Filter(g Graph, fromAttrsVals, toAttrsVals map[Attribute]string) Graph {
	out := g.New()

	fromAttrs := NodeAttributes{}
	for k, v := range fromAttrsVals {
		id := g.AttrID(k, v)
		if v != "" && id == -1 {
			return out
		}
		fromAttrs[k] = id
	}

	toAttrs := NodeAttributes{}
	for k, v := range toAttrsVals {
		id := g.AttrID(k, v)
		if v != "" && id == -1 {
			return out
		}
		toAttrs[k] = id
	}

	for _, edge := range g.Edges {
		fromNode := g.Nodes[edge.FromNode]
		toNode := g.Nodes[edge.ToNode]

		if !attributesMatch(fromAttrs, fromNode.Attrs) {
			continue
		}

		if !attributesMatch(toAttrs, toNode.Attrs) {
			continue
		}

		fromID := out.AddNode(fromNode)
		toID := out.AddNode(toNode)
		out.AddEdge(Edge{
			FromNode: fromID,
			ToNode:   toID,
			Attrs:    edge.Attrs,
		})
	}

	return out
}

func attributesMatch(c, t NodeAttributes) bool {
	for k, v := range c {
		nodeV, ok := t[k]
		match := (!ok && v == -1) || nodeV == v
		if !match {
			return false
		}
	}
	return true
}

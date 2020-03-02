package graph

type GraphVizNode struct {
	ID string `json:"id,omitempty"`
}

type GraphVizLink struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
}

type GraphViz struct {
	Nodes []GraphVizNode `json:"nodes,omitempty"`
	Links []GraphVizLink `json:"links,omitempty"`
}

func ToGraphViz(g Graph, dim Attribute) GraphViz {
	out := GraphViz{}

	done := map[string]bool{}

	for _, l := range g.Edges {
		from := g.NodeAttr(l.FromNode, dim)
		if from == "" {
			from = "N/A"
		}
		to := g.NodeAttr(l.ToNode, dim)
		if to == "" {
			to = "N/A"
		}

		if !done[from] {
			out.Nodes = append(out.Nodes, GraphVizNode{from})
			done[from] = true
		}
		if !done[to] {
			out.Nodes = append(out.Nodes, GraphVizNode{to})
			done[to] = true
		}

		out.Links = append(out.Links, GraphVizLink{
			Source: from,
			Target: to,
		})
	}

	return out
}

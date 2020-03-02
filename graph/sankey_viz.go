package graph

import "fmt"

type Sankey struct {
	Nodes []SankeyNode `json:"nodes"`
	Links []SankeyLink `json:"links"`
}

type SankeyNode struct {
	Name string `json:"name"`
}

type SankeyLink struct {
	Value float64 `json:"value"`
	Nodes []int   `json:"nodes"`
}

func ToSankey(g Graph, fromDims, toDims []Attribute, value Attribute) Sankey {
	out := Sankey{}

	nodeMap := map[string]int{}
	nodeIdx := -1
	for _, edge := range g.Edges {
		sanLink := SankeyLink{
			Value: edge.Attrs[value],
		}

		for _, d := range fromDims {
			v := g.NodeAttr(edge.FromNode, d)
			if v == "" {
				v = "N/A"
			}

			dk := fmt.Sprintf("src-%d=%s", d, v)
			if _, ok := nodeMap[dk]; !ok {
				nodeIdx++
				out.Nodes = append(out.Nodes, SankeyNode{Name: v})
				nodeMap[dk] = nodeIdx
			}

			sanLink.Nodes = append(sanLink.Nodes, nodeMap[dk])
		}

		for _, d := range toDims {
			v := g.NodeAttr(edge.ToNode, d)
			if v == "" {
				v = "N/A"
			}

			dk := fmt.Sprintf("dst-%d=%s", d, v)
			if _, ok := nodeMap[dk]; !ok {
				nodeIdx++
				out.Nodes = append(out.Nodes, SankeyNode{Name: v})
				nodeMap[dk] = nodeIdx
			}

			sanLink.Nodes = append(sanLink.Nodes, nodeMap[dk])
		}

		out.Links = append(out.Links, sanLink)
	}

	return out

}

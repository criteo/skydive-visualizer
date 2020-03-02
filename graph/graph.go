package graph

import "sort"

type Graph struct {
	AttrsDescs []AttributeDesc
	Nodes      []Node
	Edges      []Edge

	nodesAttrsIdxC   int
	nodesAttrsIdx    map[Attribute]map[string]int
	nodesAttrsIdxRev map[Attribute]map[int]string
}

func (g Graph) New() Graph {
	return Graph{
		AttrsDescs:       g.AttrsDescs,
		nodesAttrsIdxC:   g.nodesAttrsIdxC,
		nodesAttrsIdx:    g.nodesAttrsIdx,
		nodesAttrsIdxRev: g.nodesAttrsIdxRev,
	}
}

func (g *Graph) AddNode(node Node) int {
	g.Nodes = append(g.Nodes, node)
	return len(g.Nodes) - 1
}

func (g *Graph) CreateNode(attrs map[Attribute]string) int {
	n := Node{Attrs: make(map[Attribute]int)}
	i := g.AddNode(n)
	for k, v := range attrs {
		g.SetNodeAttr(i, k, v)
	}
	return i
}

func (g *Graph) SetNodeAttr(node int, attr Attribute, value string) {
	if g.nodesAttrsIdx == nil {
		g.nodesAttrsIdx = map[Attribute]map[string]int{}
	}
	if g.nodesAttrsIdxRev == nil {
		g.nodesAttrsIdxRev = map[Attribute]map[int]string{}
	}
	if g.nodesAttrsIdx[attr] == nil {
		g.nodesAttrsIdx[attr] = map[string]int{}
	}
	if g.nodesAttrsIdxRev[attr] == nil {
		g.nodesAttrsIdxRev[attr] = map[int]string{}
	}

	id, ok := g.nodesAttrsIdx[attr][value]
	if !ok {
		g.nodesAttrsIdxC++
		g.nodesAttrsIdx[attr][value] = g.nodesAttrsIdxC
		g.nodesAttrsIdxRev[attr][g.nodesAttrsIdxC] = value
		id = g.nodesAttrsIdxC
	}

	g.Nodes[node].Attrs[attr] = id
}

func (g *Graph) NodeAttr(node int, attr Attribute) string {
	if g.nodesAttrsIdxRev == nil {
		return ""
	}
	if g.nodesAttrsIdxRev[attr] == nil {
		return ""
	}
	return g.nodesAttrsIdxRev[attr][g.Nodes[node].Attrs[attr]]
}

func (g *Graph) AttrID(attr Attribute, value string) int {
	if g.nodesAttrsIdx == nil {
		return -1
	}

	if g.nodesAttrsIdx[attr] == nil {
		return -1
	}

	if v, ok := g.nodesAttrsIdx[attr][value]; ok {
		return v
	}

	return -1
}

func (g *Graph) AttrValues(attr Attribute) []string {
	if g.nodesAttrsIdx == nil {
		return []string{}
	}

	if g.nodesAttrsIdx[attr] == nil {
		return []string{}
	}

	res := []string{}
	for k := range g.nodesAttrsIdx[attr] {
		if k != "" {
			res = append(res, k)
		}
	}

	sort.Strings(res)
	return res
}

func (g *Graph) AddEdge(l Edge) int {
	g.Edges = append(g.Edges, l)
	return len(g.Edges) - 1
}

func (g *Graph) AddAttribute(id Attribute, name string, t AttributeType) {
	for _, a := range g.AttrsDescs {
		if a.ID == id {
			return
		}
	}
	g.AttrsDescs = append(g.AttrsDescs, AttributeDesc{
		Name: name,
		ID:   id,
		Type: t,
	})

	sort.Slice(g.AttrsDescs, func(i, j int) bool {
		return g.AttrsDescs[i].ID < g.AttrsDescs[j].ID
	})
}

func (g *Graph) Attribute(id Attribute) AttributeDesc {
	for _, a := range g.AttrsDescs {
		if a.ID == id {
			return a
		}
	}
	return AttributeDesc{}
}

type Node struct {
	Attrs NodeAttributes
}

type Edge struct {
	FromNode int
	ToNode   int
	Attrs    EdgeAttributes
}

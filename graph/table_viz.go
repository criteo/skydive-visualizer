package graph

type Table struct {
	Header []string        `json:"header,omitempty"`
	Rows   [][]interface{} `json:"rows,omitempty"`
}

func ToTable(g Graph, fromDims, toDims []Attribute, value Attribute) Table {
	out := Table{}

	for _, d := range fromDims {
		out.Header = append(out.Header, "Source: "+g.Attribute(d).Name)
	}
	for _, d := range toDims {
		out.Header = append(out.Header, "Destination: "+g.Attribute(d).Name)
	}
	out.Header = append(out.Header, g.Attribute(value).Name)

	for _, link := range g.Edges {
		row := []interface{}{}

		for _, d := range fromDims {
			v := g.NodeAttr(link.FromNode, d)
			row = append(row, v)
		}

		for _, d := range toDims {
			v := g.NodeAttr(link.ToNode, d)
			row = append(row, v)
		}

		row = append(row, link.Attrs[value])

		out.Rows = append(out.Rows, row)
	}

	return out
}

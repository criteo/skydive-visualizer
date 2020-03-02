package graph

type Attribute int

type NodeAttributes map[Attribute]int

func (a NodeAttributes) Equals(o NodeAttributes) bool {
	if len(a) != len(o) {
		return false
	}

	for k, v := range a {
		if o[k] != v {
			return false
		}
	}

	return true
}

type EdgeAttributes map[Attribute]float64

func (a EdgeAttributes) Copy() EdgeAttributes {
	res := make(EdgeAttributes, len(a))
	for k, v := range a {
		res[k] = v
	}
	return res
}

type AttributeType string

const (
	AttributeTypeNode AttributeType = "node"
	AttributeTypeEdge AttributeType = "edge"
)

type AttributeDesc struct {
	Name string        `json:"name,omitempty"`
	ID   Attribute     `json:"id,omitempty"`
	Type AttributeType `json:"type,omitempty"`
}

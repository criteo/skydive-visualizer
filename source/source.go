package source

import "network/skydive-visualizer-go/graph"

type Source interface {
	Fetch() (graph.Graph, error)
}

//go:generate bash -c "./generate > port_list.go"

package ports

import (
	"network/skydive-visualizer-go/graph"
	"network/skydive-visualizer-go/source"
	"network/skydive-visualizer-go/source/skydive"

	"github.com/prometheus/common/log"
)

type PortsSource struct {
	inner source.Source
}

func New(inner source.Source) *PortsSource {
	return &PortsSource{inner}
}

func (s *PortsSource) Fetch() (graph.Graph, error) {
	log.Info("ports source: starting")
	defer log.Info("ports source: done")

	g, err := s.inner.Fetch()
	if err != nil {
		return g, err
	}

	g.AddAttribute(AttrPortName, "Port name", graph.AttributeTypeNode)

	for i := range g.Nodes {
		port := g.NodeAttr(i, skydive.AttributePortNum)
		g.SetNodeAttr(i, AttrPortName, ports[port])
	}

	return g, nil
}

package dns

import (
	"net"
	"network/skydive-visualizer-go/graph"
	"network/skydive-visualizer-go/source"
	"network/skydive-visualizer-go/source/skydive"
	"strings"

	log "github.com/sirupsen/logrus"
)

type DNS struct {
	inner source.Source
}

func New(inner source.Source) *DNS {
	return &DNS{inner}
}

func (s *DNS) Fetch() (graph.Graph, error) {
	log.Info("dns source: starting")
	defer log.Info("dns source: done")

	g, err := s.inner.Fetch()
	if err != nil {
		return graph.Graph{}, err
	}

	g.AddAttribute(AttributeHostname, "Hostname", graph.AttributeTypeNode)

	cache := map[string]string{}
	for k := range g.Nodes {
		ip := g.NodeAttr(k, skydive.AttributeIPAddr)
		if ip == "" {
			continue
		}
		var result string
		if v, ok := cache[ip]; ok {
			result = v
		} else {
			names, err := net.LookupAddr(ip)
			if err != nil {
				log.Warn(err.Error())
				continue
			}
			if len(names) == 0 {
				continue
			}
			result = strings.TrimSuffix(names[0], ".")
			cache[ip] = result
		}

		g.SetNodeAttr(k, AttributeHostname, result)
	}

	return g, nil
}

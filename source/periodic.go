package source

import (
	"network/skydive-visualizer-go/graph"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Periodic struct {
	inner Source
	freq  time.Duration
	ready sync.WaitGroup
	g     graph.Graph
}

func NewPeriodic(inner Source, freq time.Duration) *Periodic {
	p := &Periodic{
		inner: inner,
		freq:  freq,
	}
	p.ready.Add(1)
	go p.run()
	return p
}

func (s *Periodic) Fetch() (graph.Graph, error) {
	s.ready.Wait()
	return s.g, nil
}

func (s *Periodic) run() {
	ready := false

	for {
		start := time.Now()
		log.Info("periodic: refreshing graph")
		g, err := s.inner.Fetch()
		if err != nil {
			log.Warnf("error fetching graph: %s", err)
			time.Sleep(s.freq)
			continue
		}
		log.Infof("periodic: graph refreshed in %s", time.Since(start))

		s.g = g
		if !ready {
			s.ready.Done()
			ready = true
		}
		time.Sleep(s.freq)
	}
}

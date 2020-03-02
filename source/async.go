package source

import "network/skydive-visualizer-go/graph"

type WillFetchResult struct {
	Graph graph.Graph
	Error error
}

func WillFetch(source Source) <-chan WillFetchResult {
	res := make(chan WillFetchResult)
	go func() {
		g, err := source.Fetch()
		res <- WillFetchResult{
			Graph: g,
			Error: err,
		}
	}()
	return res
}

package skydive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type gremlinQuery struct {
	GremlinQuery string
}

type Skydive struct {
	url    string
	client *http.Client
}

func New(url string) *Skydive {
	return &Skydive{
		url: url,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (s *Skydive) LookupNodes(ctx context.Context, query string) ([]Node, error) {
	reqBody := &bytes.Buffer{}
	json.NewEncoder(reqBody).Encode(gremlinQuery{
		GremlinQuery: query,
	})

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/topology", s.url), reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "skydive: lookup nodes")
	}
	req = req.WithContext(ctx)

	res, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "skydive: lookup nodes")
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("skydive: lookup nodes: bad status code %d", res.StatusCode)
	}

	var nodes NodeResponse

	err = json.NewDecoder(res.Body).Decode(&nodes)
	if err != nil {
		return nil, errors.Wrap(err, "skydive: lookup nodes")
	}

	return nodes.Nodes, nil
}

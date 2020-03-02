package server

import (
	"encoding/json"
	"net/http"
	"network/skydive-visualizer-go/graph"

	"github.com/julienschmidt/httprouter"
)

type graphMultiDimReq struct {
	Volume     graph.Attribute `json:"volume,omitempty"`
	MaxValues  int             `json:"maxValues,omitempty"`
	Dimensions struct {
		Src []graph.Attribute `json:"src,omitempty"`
		Dst []graph.Attribute `json:"dst,omitempty"`
	} `json:"dimensions,omitempty"`
	Filters struct {
		Src map[graph.Attribute]string `json:"src,omitempty"`
		Dst map[graph.Attribute]string `json:"dst,omitempty"`
	} `json:"filters,omitempty"`
}

type graphSingleDimReq struct {
	Volume    graph.Attribute `json:"volume,omitempty"`
	MaxValues int             `json:"maxValues,omitempty"`
	Dimension graph.Attribute
	Filters   struct {
		Src map[graph.Attribute]string `json:"src,omitempty"`
		Dst map[graph.Attribute]string `json:"dst,omitempty"`
	} `json:"filters,omitempty"`
}

func (s *Server) graphSankey(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := graphMultiDimReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	g, err := s.multiDimGraph(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	res := graph.ToSankey(g, req.Dimensions.Src, req.Dimensions.Dst, req.Volume)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (s *Server) graphTable(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := graphMultiDimReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	g, err := s.multiDimGraph(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	res := graph.ToTable(g, req.Dimensions.Src, req.Dimensions.Dst, req.Volume)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (s *Server) graphGraph(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	req := graphSingleDimReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	g, err := s.singleDimGraph(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	res := graph.ToGraphViz(g, req.Dimension)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (s *Server) multiDimGraph(req graphMultiDimReq) (graph.Graph, error) {
	g, err := s.skydive.Fetch()
	if err != nil {
		return g, err
	}

	g = graph.Filter(g, req.Filters.Src, req.Filters.Dst)
	g = graph.Group(g, req.Dimensions.Src, req.Dimensions.Dst)
	g = graph.Top(g, req.MaxValues, req.Volume)

	return g, nil
}

func (s *Server) singleDimGraph(req graphSingleDimReq) (graph.Graph, error) {
	g, err := s.skydive.Fetch()
	if err != nil {
		return g, err
	}

	g = graph.Filter(g, req.Filters.Src, req.Filters.Dst)
	g = graph.Group(g, []graph.Attribute{req.Dimension}, []graph.Attribute{req.Dimension})
	g = graph.Top(g, req.MaxValues, req.Volume)

	return g, nil
}

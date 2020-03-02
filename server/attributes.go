package server

import (
	"encoding/json"
	"net/http"
	"network/skydive-visualizer-go/graph"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) attributes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	g, err := s.skydive.Fetch()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g.AttrsDescs)
}

func (s *Server) attributeValues(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	g, err := s.skydive.Fetch()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	attr, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g.AttrValues(graph.Attribute(attr)))
}

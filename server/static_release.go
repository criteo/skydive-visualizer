// +build release

package server

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"

	_ "network/skydive-visualizer-go/server/public"
)

func (s *Server) serveStatic() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	s.router.Handler("GET", "/web/*filepath", http.StripPrefix("/web/", http.FileServer(statikFS)))
}

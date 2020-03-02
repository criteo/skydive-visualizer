package server

import (
	"net/http"
	"network/skydive-visualizer-go/source"
	"os"
	"path"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	listen  string
	router  *httprouter.Router
	skydive source.Source
}

func New(listen string, skydive source.Source) *Server {
	return &Server{
		listen:  listen,
		router:  httprouter.New(),
		skydive: skydive,
	}
}

func (s *Server) Serve() error {
	cwd, _ := os.Getwd()
	path := path.Join(cwd, path.Dir(os.Args[0]), "public")
	http.FileServer(http.Dir(path))

	s.serveStatic()

	s.router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.Redirect(w, r, "/web", http.StatusMovedPermanently)
	})

	s.router.GET("/status", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("OK"))
	})

	s.router.GET("/attributes", s.attributes)
	s.router.GET("/attributes/:id/values", s.attributeValues)
	s.router.POST("/graph/sankey", s.graphSankey)
	s.router.POST("/graph/table", s.graphTable)
	s.router.POST("/graph/graph", s.graphGraph)

	log.Infof("listening on %s", s.listen)
	return http.ListenAndServe(s.listen, s.router)
}

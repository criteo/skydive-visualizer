// +build !release

package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) serveStatic() {
	remote, err := url.Parse("http://127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	s.router.GET("/web/*filepath", handler(proxy))
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		p.ServeHTTP(w, r)
	}
}

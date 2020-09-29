package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chenminhua/httpcache/cache"
)

type Server struct {
	cache.Cache
}

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	b, e := json.Marshal(h.GetStat())
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":12345", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

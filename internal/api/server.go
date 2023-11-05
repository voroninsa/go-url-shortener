package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/voroninsa/go-url-shortener/internal/controllers"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/{short_url}", s.getHandler()).Methods(http.MethodGet)
	s.HandleFunc("/shorten-url", s.postHandler()).Methods(http.MethodPost)
}

func (s *Server) getHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := controllers.InstantiateController(r)
		url, err := c.GetUrl()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else if url == "" {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "Not found short URL.")
		}
		io.WriteString(w, url)
	}
}

func (s *Server) postHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := controllers.InstantiateController(r)
		short_url, err := c.CreateShortUrl()

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		io.WriteString(w, short_url)
	}
}

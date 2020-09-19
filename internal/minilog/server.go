package minilog

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/remicaumette/minilog/internal/minilog/handler"
	"github.com/remicaumette/minilog/internal/minilog/store"
	"net/http"
)

type Server struct {
	store      *store.Store
	httpServer *http.Server
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	dbStore, err := store.New("./db")
	if err != nil {
		return err
	}
	s.store = dbStore

	router := chi.NewRouter()
	router.Get("/query", handler.HandleQuery(dbStore))
	router.Post("/ingest", handler.HandleIngest(dbStore))

	s.httpServer = &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	if err := s.store.Close(); err != nil {
		return err
	}

	return nil
}

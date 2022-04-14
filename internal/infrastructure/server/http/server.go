package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

type server struct {
	server *http.Server
	log    log.Logger
}

func New(handler http.Handler, port string, log log.Logger) *server {
	return &server{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
			ReadTimeout: 7 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
		log: log,
	}
}

func (s *server) ListenAndServe() {
	go func() {
		s.log.Printf("Server is runnng on %s", s.server.Addr)

		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Printf("error on Listen and Serve: %v", err)
		}
	}()
}

func (s *server) Shutdown() {
	s.log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		s.log.Printf("could not shutdown server in 60s: %v", err)
		return
	}

	s.log.Println("server gracefully stopped")
}
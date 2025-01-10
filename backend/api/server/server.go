package server

import (
	"log"
	"net/http"
)

type handler interface {
	RegisterRoutes() http.Handler
}

type server struct {
	listenAddr string
	handler    handler
}

func New(listenAddr string, h handler) *server {
	return &server{
		listenAddr: listenAddr,
		handler:    h,
	}
}

func (s *server) Run() error {
	router := s.handler.RegisterRoutes()

	log.Println("server is running on port", s.listenAddr, "~")

	if err := http.ListenAndServe(s.listenAddr, router); err != nil {
		return err
	}

	return nil
}

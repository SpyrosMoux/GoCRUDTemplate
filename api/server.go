package api

import (
	"golangSampleCRUD/storage"
	"net/http"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserByID)
	http.HandleFunc("/user/id", s.handleDeleteUserByID)
	return http.ListenAndServe(s.listenAddr, nil)
}

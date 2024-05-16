package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("userId")
		w.Write([]byte("User ID: " + userId))
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}
	log.Printf("API server started on %s", s.addr)
	return server.ListenAndServe()
}

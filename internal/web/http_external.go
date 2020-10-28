package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	address string
}

func New(address string) *Server {
	return &Server{
		router:  mux.NewRouter(),
		address: address,
	}
}

func (s *Server) Start() error {
	var err error

	s.NewRouter()

	log.Println("Server started: ", s.address)
	if err = http.ListenAndServe(s.address, s.router); err != nil {
		return err
	}

	return nil
}

func (s *Server) NewRouter() {
	s.router.HandleFunc("/parse_by_nickname", s.parseByNicknameHandler).Methods("POST")
}

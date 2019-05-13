package urlshort

import (
	"log"
	"os"
)

type ServerConfig struct {
	Store       *Store
	RedirectReg *RedirectRegistry
}

type Server struct {
	store       *Store
	redirectReg *RedirectRegistry
	Router      *Router
	log         *log.Logger
}

func MakeServer(conf *ServerConfig) *Server {
	return &Server{
		store:       conf.Store,
		redirectReg: conf.RedirectReg,
		Router:      &Router{},
		log:         log.New(os.Stdout, "", log.LstdFlags),
	}
}

var CORE_ENDPOINTS map[string]string = map[string]string{
	"api":    "/api/",
	"health": "/health/",
}

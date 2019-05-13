package urlshort

import (
	"fmt"
	"net/http"
)

func handleAPIRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/api/shorten\n")
	}
}

func handleAPIRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// res := engine.AddRedirectFromJSON(r.body)
		// if already registered, return error
		// if successfully registered, return success
		fmt.Fprintf(w, "Not implemented")
	}
}

func (server *Server) addAPIRoutes() {
	server.Router.coreHandler.Handle("/api", handleAPIRoot())
	server.Router.coreHandler.Handle("/api/shorten", handleAPIRegister())
}

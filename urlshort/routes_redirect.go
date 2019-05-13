package urlshort

import (
	"fmt"
	"net/http"
)

func handleMissing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Invalid path")
}

// Handle redirects from server's redirect registry; fallback on servers previous handler
func (server *Server) addRegistryRedirects() {
	prevHandler := server.Router.redirectsHandler
	server.Router.redirectsHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if redirect, ok := (*server.redirectReg).Get(r.URL.Path); ok {
				http.Redirect(w, r, redirect.Dest, 300)
			} else {
				prevHandler.ServeHTTP(w, r)
			}
		})
}

// Handle redirects from server's store; fallback on servers previous handler
func (server *Server) addStoreRedirects() {
	prevHandler := server.Router.redirectsHandler
	server.Router.redirectsHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			redirect, err := (*server.store).Get(r.URL.Path)
			if redirect != nil && err == nil {
				http.Redirect(w, r, redirect.Dest, 300)
			} else {
				prevHandler.ServeHTTP(w, r)
			}
		})
}

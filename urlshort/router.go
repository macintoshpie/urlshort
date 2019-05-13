package urlshort

import (
	"net/http"
	"strings"
)

// Router for server
type Router struct {
	coreHandler      *http.ServeMux
	redirectsHandler http.Handler
}

// Attempts to handle requests with core endpoints, then redirect endpoints
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, endpoint := range CORE_ENDPOINTS {
		if strings.HasPrefix(r.URL.Path, endpoint) {
			router.coreHandler.ServeHTTP(w, r)
			return
		}
	}
	router.redirectsHandler.ServeHTTP(w, r)
}

// Initializes routes for the server
func (server *Server) SetupRouter() {
	server.Router.coreHandler = http.NewServeMux()
	server.Router.redirectsHandler = http.HandlerFunc(handleMissing)

	// Handle core routes
	server.addAPIRoutes()

	// Handle redirect routes
	// Searches for redirects in this order:
	//   1. redirect registry
	//   2. store
	//   3. missing redirect
	if *server.store != nil {
		server.addStoreRedirects()
	}
	server.addRegistryRedirects()
}

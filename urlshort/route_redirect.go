package urlshort

import "net/http"

// Update router to handle registered redirects
func applyRedirectRegistry(redirectReg *RedirectRegistry, mux *http.ServeMux) {
	for _, redirect := range redirectReg.Redirects {
		rh := http.RedirectHandler(redirect.Dest, 307)
		mux.Handle(redirect.Src, rh)
	}
}

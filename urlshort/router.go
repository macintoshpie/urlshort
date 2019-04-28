package urlshort

import (
	"net/http"
)

func MakeServeMux(redirectReg *RedirectRegistry) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	applyRedirectRegistry(redirectReg, mux)
	return mux
}

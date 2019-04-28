package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/macintoshpie/urlshort/urlshort"
)

func main() {
	var jsonRedirectsEnv string = os.Getenv("JSON_REDIRECTS_PATH")
	var port *uint = flag.Uint("port", 8080, "port on which to expose the API")
	var jsonRedirects *string = flag.String(
		"json",
		jsonRedirectsEnv,
		"path to json file for preconfiguring redirects",
	)
	flag.Parse()

	if *jsonRedirects == "" {
		fmt.Println("INFO: no $JSON_REDIRECTS_PATH or --json provided; router will initialize with no redirects")
	}

	// create a redirect registry
	redirectReg := urlshort.NewRedirectRegistry()
	err := redirectReg.AddFromJSON(*jsonRedirects)
	if err != nil {
		fmt.Printf("WARNING: failed to load redirects from json: %s\n", err)
	}

	// create request router
	mux := urlshort.MakeServeMux(redirectReg)

	httpLogger := log.New(os.Stdout, "", log.LstdFlags)
	addr := fmt.Sprintf(":%d", *port)
	httpServer := &http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}
	httpLogger.Println(fmt.Sprintf("INFO: Starting server at %s", addr))
	httpLogger.Fatal(httpServer.ListenAndServe())
}

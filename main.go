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
	var dbName *string = flag.String("dbname", "", "Database name")
	flag.Parse()

	redirectReg := urlshort.NewRedirectRegistry()
	if *jsonRedirects == "" {
		fmt.Println("INFO: no $JSON_REDIRECTS_PATH or --json provided; router will initialize with no redirects")
	} else {
		// create a redirect registry
		err := redirectReg.AddFromJSON(*jsonRedirects)
		if err != nil {
			fmt.Printf("WARNING: failed to load redirects from json: %s\n", err)
		}
	}

	var store urlshort.Store
	var err error
	if *dbName == "" {
		store = nil
		fmt.Println("INFO: no --dbname provided")
	} else {
		// create store for redirects
		store, err = urlshort.MakeRDBStore(fmt.Sprintf("dbname=%v sslmode=disable", *dbName))
		if err != nil {
			fmt.Printf("ERROR: failed to create database: %s\n", err)
			return
		}
	}

	// create server
	serverConfig := urlshort.ServerConfig{
		Store:       &store,
		RedirectReg: redirectReg,
	}
	shortServer := urlshort.MakeServer(&serverConfig)
	shortServer.SetupRouter()

	// r, err := store.Get("/fbk")
	// if err != nil {
	// 	fmt.Printf("Failed to get: %v\n", err)
	// } else if r == nil {
	// 	fbkRedir := urlshort.Redirect{Src: "/fbk", Dest: "http://www.facebook.com", Ttl: 100}
	// 	err = store.Add(&fbkRedir)
	// 	if err != nil {
	// 		fmt.Printf("Failed to add redirect; %v", err)
	// 	}
	// } else {
	// 	fmt.Printf("%s", r.Dest)
	// }

	httpLogger := log.New(os.Stdout, "", log.LstdFlags)
	addr := fmt.Sprintf(":%d", *port)
	httpServer := &http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      shortServer.Router,
	}
	httpLogger.Println(fmt.Sprintf("INFO: Starting server at %s", addr))
	httpLogger.Fatal(httpServer.ListenAndServe())
}

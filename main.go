package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	numWorker = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr  = flag.String("http", "127.0.0.1:8000", "Address to listen")
)

func main() {
	flag.Parse()
	// start dispatcher
	StartDispatcher(*numWorker)
	// register Collector
	http.HandleFunc("/wort", Collector)
	// start http server
	log.Println("HTTP server listening on ", *HTTPAddr)
	log.Fatal(http.ListenAndServe(*HTTPAddr, nil))
}

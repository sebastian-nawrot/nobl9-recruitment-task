package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sebastian-nawrot/nobl9-recruitment-task/api"
)

func main() {
	port := flag.Int("port", 10000, "Listening port of our API")
	flag.Parse()

	fmt.Printf("Listening on port: %d\n", *port)

	http.HandleFunc("/random/mean", api.RandomMeanEndpoint)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

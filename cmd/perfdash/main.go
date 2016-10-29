package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"chain/env"
)

var (
	addr    = env.String("LISTEN", ":8080")
	baseURL = env.String("BASE_URL", "http://localhost:1999/") // see html.go
)

func main() {
	env.Parse()

	if len(os.Args) > 1 && os.Args[1] == "cron" {
		cron()
		return
	} else if len(os.Args) > 1 {
		usage()
		os.Exit(2)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/histogram.png", histogram)
	log.Fatalln(http.ListenAndServe(*addr, nil))
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: perfdash [cron]")
}

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"chain/env"
)

var (
	addr    = env.String("LISTEN", ":8080")
	baseURL = env.String("BASE_URL", "http://localhost:1999/")
)

var sizes = []time.Duration{
	time.Minute, // first size must == bucket size
	time.Hour,   // each size must be an integral multiple of previous
	24 * time.Hour,
}

func main() {
	env.Parse()

	if len(os.Args) > 1 && os.Args[1] == "cron" {
		cron()
		return
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/histogram.png", histogram)
	log.Fatalln(http.ListenAndServe(*addr, nil))
}

func cron() {
	for t := range time.Tick(period) {
		rotateSchema(t)
		fetchMetrics()
	}
}

func rotateSchema(t time.Time) {
	for _, size := range sizes {
		rotateSchemaSize(t, size)
	}
}

func rotateSchemaSize(t time.Time, size time.Duration) {
	var zero time.Time
	if t.Sub(zero)%size == 0 {
	}
}

func fetchMetrics() {
}

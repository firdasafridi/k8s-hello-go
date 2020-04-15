package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var port string
var count int

func init() {
	flag.StringVar(&port, "port", ":8080", "Please define port, if not it will default :8080")
	flag.IntVar(&count, "count", 1000, "Define value at count logic")
	flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	for i := 1; i <= count; i++ {

	}
	fmt.Fprintf(w, "%s %f", r.URL.Path[1:], time.Since(start).Seconds())
}

func main() {
	http.HandleFunc("/", handler)

	log.Printf("The application run on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

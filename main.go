package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    port := "8080"
    if p := os.Getenv("PORT"); p != "" {
        port = p
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello from myapp!"))
    })

    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(200)
        w.Write([]byte("ok"))
    })

    // Prometheus metrics endpoint
    http.Handle("/metrics", promhttp.Handler())

    addr := fmt.Sprintf(":%s", port)
    log.Printf("listening on %s", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}

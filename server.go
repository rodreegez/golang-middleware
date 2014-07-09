package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.Handle("/hello", helloHandler())
	logger := log.New(os.Stdout, "", 0)
	http.Handle("/hello-log", loggingMiddleware(logger, helloHandler()))
	http.ListenAndServe(":5000", nil)
}

func helloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("well hellooo"))
	})
}

func loggingMiddleware(l *log.Logger, n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		n.ServeHTTP(w, r)
		l.Printf("Req took %s to complete", time.Since(start))
	})
}

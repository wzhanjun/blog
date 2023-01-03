package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var started time.Time

func main() {
	started = time.Now()
	http.HandleFunc("/", hello)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()

	dbUrl := os.Getenv("DB_URL")

	dbPassword := os.Getenv("DB_PASSWORD")

	io.WriteString(w, fmt.Sprintf("[v5] Hello, Kubernetes!, From host:%s, Get Database Connect URL:%s, password:%s ", host, dbUrl, dbPassword))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// duration := time.Since(started)
	// if duration.Seconds() > 15 {
	// 	w.WriteHeader(500)
	// 	w.Write([]byte(fmt.Sprintf("error:%v", duration.Seconds())))
	// } else {
	// 	w.WriteHeader(200)
	// 	w.Write([]byte("ok"))
	// }
	w.WriteHeader(500)
}

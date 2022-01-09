package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	os.Setenv("VERSION", "0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	clientIP := getCurrentIP(r)
	httpCode := http.StatusOK
	log.Printf("clientIP: %s \n", clientIP)
	log.Printf("httpCode: %d \n", httpCode)

	w.Header().Set("clientIP", clientIP)
	w.Header().Set("httpCode", string(rune(httpCode)))

}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("healthz ok")

}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-REAL-IP")
	if ip == "" {
		// 格式： IP:port
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("healthz/", healthz)
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal("start server failed, %s\n", err.Error())
	}
}

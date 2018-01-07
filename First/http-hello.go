package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Access IP and Port is : %s", r.RemoteAddr)
}
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! %s ", r.URL.Path)
}

func main() {
	http.HandleFunc("/path/", handler1)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8800", nil)
}

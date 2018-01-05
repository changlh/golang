package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8800", nil)
}

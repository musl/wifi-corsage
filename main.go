// vim: set ft=go sw=2 ts=2 :

package main

import (
	"fmt"
	"net/http"
)

var status = 0

func route_index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(".")
	fmt.Fprintf(w, "%d", status)
}

func main() {
	http.HandleFunc("/", route_index)
	fmt.Println("Listening on 0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

// vim: set ft=go sw=2 ts=2 :

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var id_chan = make(chan int, 1)
var code int
var code_mutex sync.Mutex

type Command struct {
	Code int
}

func finish(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, message)
}

func logDuration(message string, start time.Time) {
	end := time.Now()
	log.Printf("%s: %v\n", message, end.Sub(start))
}

func route_code(w http.ResponseWriter, r *http.Request) {
	id := <-id_chan
	start := time.Now()

	defer r.Body.Close()
	defer logDuration(fmt.Sprintf("%08d %s", id, r.URL.Path), start)

	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
		c := Command{}
		err := decoder.Decode(&c)
		if err != nil {
			log.Printf("Unprocessable Entity: %#v", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		code_mutex.Lock()
		code = c.Code
		log.Printf("Code set to: %s\n", c)
		code_mutex.Unlock()

		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method == "GET" {
		code_mutex.Lock()
		c := code
		log.Printf("Code read as: %s\n", c)
		code_mutex.Unlock()

		fmt.Fprintf(w, "%d", c)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	bind_addr := "0.0.0.0:8080"

	go func() {
		for i := 0; ; i++ {
			id_chan <- i
		}
	}()

	http.Handle("/", fs)
	http.HandleFunc("/code", route_code)

	log.Printf("Listening on %s\n", bind_addr)
	log.Fatal(http.ListenAndServe(bind_addr, nil))
}

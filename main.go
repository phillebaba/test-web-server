package main

import (
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Lainecloud is up and running"

	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", index)

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

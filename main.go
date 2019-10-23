package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "Hello World"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})

	log.Println("Serving web page")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	conso

	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":"+string(port), nil); err != nil {
		panic(err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	crawl()
	server()
}

func server() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Println("go server listening on PORT", port, "!")

	http.HandleFunc("/", respondHello)
	http.HandleFunc("/api/feeds/", getFeeds)
	http.HandleFunc("/api/feeds/reddit/", getReddit)
	http.HandleFunc("/api/feeds/wikipedia/", getWikipedia)
	http.HandleFunc("/api/feeds/github/", getGithub)
	http.HandleFunc("/api/feeds/twitter/", getTwitter)

	if err := http.ListenAndServe(":"+string(port), nil); err != nil {
		panic(err)
	}
}

func respondHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

// TODO: time for a closure methinks
func getFeeds(w http.ResponseWriter, r *http.Request) {
	jsonString := readJson("./tmp/all.json")
	w.Write(jsonString)
}

func getReddit(w http.ResponseWriter, r *http.Request) {
	jsonString := readJson("./tmp/reddit.json")
	w.Write(jsonString)
}

func getWikipedia(w http.ResponseWriter, r *http.Request) {
	jsonString := readJson("./tmp/wikipedia.json")
	w.Write(jsonString)
}

func getGithub(w http.ResponseWriter, r *http.Request) {
	jsonString := readJson("./tmp/github.json")
	w.Write(jsonString)
}

func getTwitter(w http.ResponseWriter, r *http.Request) {
	jsonString := readJson("./tmp/twitter.json")
	w.Write(jsonString)
}

// TODO: Hackernews?
// https://github.com/HackerNews/API

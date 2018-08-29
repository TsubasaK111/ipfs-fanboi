package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func check(e error, msg string) {
	if e != nil {
		log.Fatalln(msg, e)
		panic(e)
	}
}

func main() {

	jsonBody := crawler()

	var response interface{}

	err := json.Unmarshal(jsonBody, &response)
	check(err, "")

	err = ioutil.WriteFile("./public/response", jsonBody, 0644)
	// file, err := os.Create("./public/dat1", jsonBody)
	// defer file.Close()
	check(err, "")

	fmt.Println(response)

	// server()
}

func crawler() []byte {
	client := &http.Client{}

	// resp, err := http.Get("https://www.reddit.com/search.json?q=ipfs")
	req, err := http.NewRequest("GET", "https://www.reddit.com/r/ipfs/search.json?q=ipfs", nil)
	check(err, "")

	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	resp, err := client.Do(req)
	check(err, "http transport error:")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err, "read error:")

	return []byte(string(body))
}

func server() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":"+string(port), nil); err != nil {
		panic(err)
	}
}

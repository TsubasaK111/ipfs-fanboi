package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func crawl() {
	redditJson := get("https://www.reddit.com/r/ipfs/search.json?q=ipfs")
	writeJson(redditJson, "./tmp/reddit.json")

	wikipediaJson := get("https://en.wikipedia.org/w/api.php?action=opensearch&search=%22ipfs%22&format=json")
	writeJson(wikipediaJson, "./tmp/wikipedia.json")

	wikiArticleJson := get("https://en.wikipedia.org/w/api.php?action=parser-migration&title=InterPlanetary_File_System&format=json&config=new")
	writeJson(wikiArticleJson, "./tmp/wikiArticle.json")

	twitterJson := twitterSearch()
	writeJson(twitterJson, "./tmp/twitter.json")

	githubJson := get("https://api.github.com/search/repositories?q=ipfs&sort=updated&order=desc")
	writeJson(githubJson, "./tmp/github.json")

	githubItems := convertGithubJson(githubJson)
	redditItems := convertRedditJson(redditJson)
	twitterItems := convertTwitterJson(twitterJson)

	aggregateItems := concatItems([]Items{
		githubItems,
		redditItems,
		twitterItems,
	})

	aggregateItems = sortItems(aggregateItems)

	itemsJson := stringifyJson(aggregateItems)
	writeJson(itemsJson, "./tmp/all.json")
}

func syncWrite(content []byte, filePath string) {
	fmt.Println(string(content))

	f, err := os.Create(filePath)
	check(err, "")

	defer fmt.Println("should close ")
	defer f.Close()

	res, err := f.Write(content)
	fmt.Printf("wrote %d bytes\n", res)

	f.Sync()
}

func get(url string) []byte {
	client := &http.Client{}

	// resp, err := http.Get(url)
	req, err := http.NewRequest("GET", url, nil)
	check(err, "")

	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	resp, err := client.Do(req)
	check(err, "http transport error:")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err, "read error:")

	return []byte(string(body))
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
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
	// other items
	itemsJson := itemsToJson(githubItems)
	writeJson(itemsJson, "/tmp/dashboard.json")
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

type Item struct {
	name        string
	description string
	url         url.URL
	created_at  time.Time
	updated_at  time.Time
	source      string
	gravity     int
}

func convertGithubJson(githubJson []byte) []Item {
	parsed := parseJson(githubJson)

	// lots of typecasting!
	rawItems := parsed["items"].([]interface{})

	rawItem := rawItems[0].(map[string]interface{})
	itemName := rawItem["name"].(string)
	itemDescription := rawItem["description"].(string)
	itemUrl := rawItem["url"].(string)

	fmt.Println(itemName, itemDescription, itemUrl)

	// for _, rawItem := range rawItems {
	// }
	item := Item{name: "Sean", description: "50"}
	item2 := Item{name: "Sean", description: "50"}

	items := []Item{item, item2}

	return items
}

func itemsToJson(items []Item) []byte {
	return []byte("wat")
}

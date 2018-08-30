package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func twitterSearch() []byte {
	client := setupClient()

	// sentTweet := sendTweet(client)
	// timelineTweets := getTimeline(client)
	// statusTweet := getStatus(client)

	searchResults := searchTweets(client)

	jsonResults := stringifyJson(searchResults)

	return jsonResults
}

func setupClient() *twitter.Client {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("TWITTER_* env vars must be set:\n",
			"consumerKey:", consumerKey,
			"consumerSecret:", consumerSecret,
			"accessToken:", accessToken,
			"accessSecret:", accessSecret,
		)
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	return client
}

func searchTweets(client *twitter.Client) *twitter.Search {
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "ipfs",
	})
	check(err, "")
	return search
}

func getStatus(client *twitter.Client) *twitter.Tweet {
	tweet, _, err := client.Statuses.Show(585613041028431872, nil)
	check(err, "")
	return tweet
}

func sendTweet(client *twitter.Client) *twitter.Tweet {
	tweet, _, err := client.Statuses.Update("just setting up my twttr", nil)
	check(err, "")
	return tweet
}

func getTimeline(client *twitter.Client) []twitter.Tweet {
	// Home Timeline
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	check(err, "")
	return tweets
}

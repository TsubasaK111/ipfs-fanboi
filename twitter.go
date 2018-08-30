package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func twitterSearch() {

	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("TWITTER_* env vars must be set:", consumerKey, consumerSecret, accessToken, accessSecret)
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	fmt.Println(tweets)

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)
	fmt.Println(tweet)
	check(err, "")

	// Status Show
	tweet, resp, err = client.Statuses.Show(585613041028431872, nil)
	check(err, "")
	fmt.Println(tweet, resp)

	// Search Tweets
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "gopher",
	})
	check(err, "")
	fmt.Println(search, resp)
}

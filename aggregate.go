package main

import (
	"net/url"
	"sort"
	"time"
)

type Item struct {
	name        string
	description string
	url         url.URL
	createdAt   time.Time
	updatedAt   time.Time
	source      string
	gravity     int
}

func sortItems(items []Item) []Item {
	sort.Slice(items, func(i, j int) bool {
		return items[i].updatedAt.Before(items[j].updatedAt)
	})
	return items
}

func convertGithubJson(githubJson []byte) []Item {
	rawJson := parseJson(githubJson)

	// lots of typecasting!
	rawItems := rawJson["items"].([]interface{})

	var items []Item

	for _, rawItem := range rawItems {
		freshItem := rawItem.(map[string]interface{})

		itemName := freshItem["name"].(string)
		kindaDescription, boolErr := freshItem["description"].(string)
		itemDescription := defaultTo(boolErr, kindaDescription, itemName)
		itemUrl, err := url.Parse(freshItem["url"].(string))
		itemCreatedAt, err := time.Parse(time.RFC3339, freshItem["created_at"].(string))
		itemUpdatedAt, err := time.Parse(time.RFC3339, freshItem["updated_at"].(string))
		stargazersCount, boolErr := freshItem["stargazers_count"].(int)
		watchersCount, boolErr := freshItem["watchers_count"].(int)
		itemGravity := stargazersCount + watchersCount
		check(err, "")

		item := Item{
			name:        itemName,
			description: itemDescription,
			url:         *itemUrl,
			createdAt:   itemCreatedAt,
			updatedAt:   itemUpdatedAt,
			source:      "github",
			gravity:     itemGravity,
		}
		items = append(items, item)
	}

	return items
}

func convertRedditJson(redditJson []byte) []Item {
	rawJson := parseJson(redditJson)

	// lots of typecasting!
	rawData := rawJson["data"].(map[string]interface{})
	rawItems := rawData["children"].([]interface{})

	var items []Item

	for _, rawItem := range rawItems {
		// rawItem := rawItems[0]
		rawItemData := rawItem.(map[string]interface{})["data"].(map[string]interface{})

		itemName := rawItemData["title"].(string)
		kindaDescription, boolErr := rawItemData["selftext"].(string)
		itemDescription := defaultTo(boolErr, kindaDescription, itemName)
		itemUrl, err := url.Parse(rawItemData["url"].(string))
		itemCreatedUtc := int64(rawItemData["created_utc"].(float64))
		itemCreatedAt := time.Unix(itemCreatedUtc, 0)
		ups, boolErr := rawItemData["ups"].(int)
		numComments, boolErr := rawItemData["num_comments"].(int)
		itemGravity := ups + numComments
		check(err, "")

		item := Item{
			name:        itemName,
			description: itemDescription,
			url:         *itemUrl,
			createdAt:   itemCreatedAt,
			updatedAt:   itemCreatedAt,
			source:      "reddit",
			gravity:     itemGravity,
		}

		items = append(items, item)
	}

	return items
}

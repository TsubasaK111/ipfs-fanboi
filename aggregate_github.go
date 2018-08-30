package main

import (
	"net/url"
	"time"
)

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
			Name:        itemName,
			Description: itemDescription,
			Url:         *itemUrl,
			CreatedAt:   itemCreatedAt,
			UpdatedAt:   itemUpdatedAt,
			Source:      "github",
			Gravity:     itemGravity,
		}
		items = append(items, item)
	}

	return items
}

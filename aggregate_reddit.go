package main

import (
	"net/url"
	"time"
)

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
		checkBool(boolErr, "")

		item := Item{
			Name:        itemName,
			Description: itemDescription,
			Url:         itemUrl.String(),
			CreatedAt:   itemCreatedAt,
			UpdatedAt:   itemCreatedAt,
			Source:      "reddit",
			Gravity:     itemGravity,
		}

		items = append(items, item)
	}

	return items
}

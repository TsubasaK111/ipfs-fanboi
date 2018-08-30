package main

import (
	"net/url"
	"sort"
	"time"
)

type Item struct {
	Name        string
	Description string
	Url         url.URL
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Source      string
	Gravity     int
}

func sortItems(items []Item) []Item {
	sort.Slice(items, func(i, j int) bool {
		return items[i].UpdatedAt.Before(items[j].UpdatedAt)
	})
	return items
}

type Items []Item

func concatItems(manyItems []Items) Items {
	var aggregateItems Items

	for _, items := range manyItems {
		aggregateItems = append(aggregateItems, items...)
	}

	return aggregateItems
}

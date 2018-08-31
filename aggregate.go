package main

import (
	"sort"
	"time"
)

type Item struct {
	Name        string
	Description string
	Url         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Source      string
	Gravity     int
}

type Items []Item

func (s Items) Len() int           { return len(s) }
func (s Items) Less(i, j int) bool { return s[i].CreatedAt.Before(s[j].CreatedAt) }
func (s Items) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func sortItems(items Items) Items {
	sort.Sort(sort.Reverse(items))
	// sort.Slice(items, func(i, j int) bool {
	// 	return items[i].UpdatedAt.Before(items[j].UpdatedAt)
	// })
	return items
}

func concatItems(manyItems []Items) Items {
	var aggregateItems Items

	for _, items := range manyItems {
		aggregateItems = append(aggregateItems, items...)
	}

	return aggregateItems
}

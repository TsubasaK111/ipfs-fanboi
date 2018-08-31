package main

func convertTwitterJson(twitterJson []byte) []Item {
	// rawJson := parseJson(twitterJson)
	// fmt.Println("rawTwitterjson:", rawJson)

	var items []Item

	item := Item{Name: "wat", Description: "twitter mockery!", Source: "twitter"}
	items = append(items, item)

	return items
}

package main

import (
	"encoding/json"
	"io/ioutil"
)

func readJson(filePath string) []byte {
	jsonString, err := ioutil.ReadFile(filePath)
	check(err, "")

	return []byte(jsonString)
}

func writeJson(wildJson []byte, filePath string) {
	err := ioutil.WriteFile(filePath, wildJson, 0644)
	// file, err := os.Create("./tmp/dat1", wildJson)
	// defer file.Close()
	check(err, "writing to json file failed")
}

func stringifyJson(items interface{}) []byte {
	jsonResults, err := json.Marshal(items)
	check(err, "WOAH")

	return jsonResults
}

func parseJson(wildJson []byte) map[string]interface{} {
	var result map[string]interface{}

	err := json.Unmarshal(wildJson, &result)
	check(err, "")

	return result
}

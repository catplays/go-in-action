package search

import (
	"encoding/json"
	"log"
	"os"
)

type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}

var DataPath = "rss/data/data.json"

func ReceiveFeed() ([]*Feed, error) {
	file, err :=  os.Open(DataPath)
	if err != nil {
		log.Printf("file %s not existed.\n", DataPath)
		return nil, err
	}
	defer file.Close()

	var feeds []*Feed
	// decode json data into feed slice
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}

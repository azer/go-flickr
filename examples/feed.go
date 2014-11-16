package main

import (
	"fmt"
	"os"
	"github.com/azer/go-flickr"
)

func main() {
	client := flickr.Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	feed, err := client.Feed("98269877@N00", 100)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d Photos in the feed. First title: %s", len(feed), feed[0].Title)
}

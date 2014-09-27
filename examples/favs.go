package main

import (
	"fmt"
	"os"
	"github.com/azer/go-flickr-client"
)

func main() {
	client := flickr.Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	favs, err := client.Favs("98269877@N00")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d Favs. First title: %s", len(favs), favs[0].Title)
	}
}

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

	following, err := client.Following("98269877@N00")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d Following. First user: %s", len(following), following[0].Title)
	}
}

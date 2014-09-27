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

	resp, err := client.Request("people.findByUsername", flickr.Params{
		"username": "azerbike",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", string(resp))
}

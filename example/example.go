package main

import (
	"github.com/azer/go-flickr-client"
	"fmt"
	"os"
)

func main () {
	client := flickr.Client(&flickr.Options{
		Key: os.Getenv("FLICKR_API_KEY"),
	})

	resp, err := client("people.findByUsername", &flickr.Params{
		"username": "azerbike",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", resp)
}

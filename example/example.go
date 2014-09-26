package main

import (
	"github.com/azer/go-flickr-client"
	"fmt"
)

func main () {
	client := flickr.Client(&flickr.Options{
		Key: "8974b9cf7bf473e056125874ad44ce0a",
	})

	resp, err := client("people.findByUsername", &flickr.Params{
		"username": "azerbike",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", resp)
}

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

	user, err := client.FindUser("azer")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s: %s", user.Id, user.Name)
	}
}

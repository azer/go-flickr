package flickr

import (
	"fmt"
	"os"
)

func Example_GettingUser() {
	client := Client(&Options{
		Key: os.Getenv("FLICKR_API_KEY"),
	})

	resp, err := client("people.findByUsername", &Params{
		"username": "azerbike",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", resp)
	// Output: map[user:map[id:39873962@N08 nsid:39873962@N08 username:map[_content:]] stat:ok]
}

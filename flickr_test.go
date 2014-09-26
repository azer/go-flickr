package flickr

import (
	"fmt"
)

func Example_GettingUser() {
	client := Client(&Options{
		Key: "8974b9cf7bf473e056125874ad44ce0a",
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

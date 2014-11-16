package flickr

import (
	"fmt"
	"os"
)

func ExampleRequest() {
	client := &Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	resp, err := client.Request("people.findByUsername", Params{
		"username": "azerbike",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", string(resp))
	// Output: {"user":{"id":"98269877@N00","nsid":"98269877@N00","username":{"_content":"azerbike"}},"stat":"ok"}
}

package flickr

import (
	"os"
	"testing"
	"fmt"
)

func TestFeed(t *testing.T) {
	client := &Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	photos, err := client.Feed("98269877@N00", 50)

	if err != nil {
		t.Error(err)
	}

	if len(photos) < 40 {
		t.Error(fmt.Printf("Less than 50 photos were returned: %d", len(photos)))
	}

	if len(photos[0].Id) == 0 {
		t.Error("First photo id is empty")
	}

	if len(photos[0].Owner) == 0 {
		t.Error("First owner is empty")
	}
}

package flickr

import (
	"os"
	"testing"
)

func TestFollowing(t *testing.T) {
	client := &Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	following, err := client.Following("98269877@N00")

	if err != nil {
		t.Error(err)
	}

	if len(following) < 90 {
		t.Error("Less than 400 following were found")
	}

	if len(following[0].Id) == 0 {
		t.Error("First user id is empty")
	}

	if len(following[0].Title) == 0 {
		t.Error("First user title is empty")
	}
}

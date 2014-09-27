package flickr

import (
	"os"
	"testing"
)

func TestFindUser(t *testing.T) {
	client := &Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	user, err := client.FindUser("azer")

	if err != nil {
		t.Error(err)
	}

	if user.Id != "98269877@N00" {
		t.Error("Invalid user id")
	}

	if user.Name != "azer" {
		t.Error("Invalid user name")
	}
}

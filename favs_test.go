package flickr

import (
	"os"
	"testing"
)

func TestFavs(t *testing.T) {
	client := &Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	favs, err := client.Favs("98269877@N00")

	if err != nil {
		t.Error(err)
	}

	if len(favs) < 90 {
		t.Error("Less than 90 favorites were created")
	}

	if len(favs[0].Id) == 0 {
		t.Error("First fav id is empty")
	}

	if len(favs[0].Owner) == 0 {
		t.Error("First owner is empty")
	}
}

package flickr

import (
	"os"
	"testing"
	"fmt"
)

func TestFavs(t *testing.T) {
	client := &Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	favs, err := client.Favs("98269877@N00")
	//favs, err := client.Favs("38322687@N06")

	if err != nil {
		t.Error(err)
	}

	if len(favs) < 50 {
		t.Error(fmt.Printf("Less than 90 favorites were created: %d", len(favs)))
	}

	if len(favs[0].Id) == 0 {
		t.Error("First fav id is empty")
	}

	if len(favs[0].Owner) == 0 {
		t.Error("First owner is empty")
	}
}

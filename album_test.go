package flickr

import (
	"fmt"
	"os"
	"testing"
)

func TestAlbum(t *testing.T) {
	client := &Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	album, err := client.Album("72157661416157179")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(album.Id)
	fmt.Println(len(album.Photos))

	if len(album.Id) == 0 {
		t.Error("album id is empty")
	}

	if len(album.Photos) != 17 {
		t.Error(fmt.Printf("Not 17 album.Photos were returned: %d", len(album.Photos)))
	}

	if len(album.Photos[0].Id) == 0 {
		t.Error("First photo id is empty")
	}

}

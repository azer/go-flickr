package flickr

import (
	"os"
	"testing"
)

func TestGetPhoto(t *testing.T) {
	client := &Client{
		Key: os.Getenv("FLICKR_API_KEY"),
	}

	photo, err := client.GetPhoto(15691826511)

	if err != nil {
		t.Error(err)
	}

	if photo.Id != 15691826511 {
		t.Error("Invalid photo")
	}

	if photo.Title != "driving to cusco" {
		t.Error("Invalid photo title")
	}

	if photo.Username != "azer" {
		t.Error("Invalid user")
	}

	if photo.UserIcon != "https://farm3.staticflickr.com/2933/buddyicons/98269877@N00_r.jpg" {
		t.Error("Invalid user icon")
	}
}

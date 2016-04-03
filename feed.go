package flickr

import "fmt"

type FeedPhoto struct {
	Id        string
	Title     string
	Owner     string
	OwnerName string
	Secret    string
	Server    string
	Farm      int
}

type FeedRaw struct {
	Photos struct {
		Photo []struct {
			Id       string
			Secret   string
			Server   string
			Farm     int
			Owner    string
			Username string
			Title    string
			IsPublic int
			IsFriend int
			IsFamily int
		}
	}
}

func (client *Client) Feed(userId string, count int) ([]FeedPhoto, error) {
	response, err := client.Request("photos.getContactsPublicPhotos", Params{
		"user_id": userId,
		"count":   fmt.Sprintf("%d", count+1),
	})

	if err != nil {
		return nil, err
	}

	raw := &FeedRaw{}
	err = Parse(response, raw)

	if err != nil {
		return nil, err
	}

	feed := []FeedPhoto{}

	for _, photo := range raw.Photos.Photo {
		feed = append(feed, FeedPhoto{
			Id:        photo.Id,
			Title:     photo.Title,
			Owner:     photo.Owner,
			OwnerName: photo.Username,
			Secret:    photo.Secret,
			Server:    photo.Server,
			Farm:      photo.Farm,
		})
	}

	return feed, nil
}

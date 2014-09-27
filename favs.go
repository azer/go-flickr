package flickr

type Fav struct {
	Id string
	Title string
	Owner string
	DateFaved string
	Farm int
	Secret string
	Server string
}

type FavsRaw struct {
	Photos struct {
		Page int
		Pages int
		PerPage int
		Total int
		Photo []struct {
			DateFaved string
			Farm int
			Id string
			IsFamily int
			IsFriend int
			IsPublic int
			Owner string
			Secret string
			Server string
			Title string
		}
	}
}

func (client *Client) Favs (userId string) ([]Fav, error) {
	response, err := client.Request("favorites.getPublicList", Params{
		"user_id": userId,
	})

	raw := &FavsRaw{}
	err = Parse(response, raw)

	if err != nil {
		return nil, err
	}

	favs := []Fav{}

	for _, photo := range raw.Photos.Photo {
		favs = append(favs, Fav{
			Id: photo.Id,
			Title: photo.Title,
			Owner: photo.Owner,
			DateFaved: photo.DateFaved,
			Farm: photo.Farm,
			Secret: photo.Secret,
			Server: photo.Server,
		})
	}

	return favs, nil
}

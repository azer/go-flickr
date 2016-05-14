package flickr

type AlbumPhoto struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Secret string `json:"secret"`
	Server string `json:"server"`
	Farm   int    `json:"farm"`
}

func (albumPhoto *AlbumPhoto) URLs() map[string]string {
	return GenerateURLs(albumPhoto.Id, albumPhoto.Farm, albumPhoto.Secret, albumPhoto.Server, "jpg")
}

type Album struct {
	Id      string        `json:"id"`
	Title   string        `json:"title"`
	Desc    string        `json:"desc"`
	Primary string        `json:"primary"`
	Owner   string        `json:"owner"`
	Photos  []*AlbumPhoto `json:"photos"`
}

func (client *Client) Album(id string) (*Album, error) {
	response, err := client.Request("photosets.getPhotos", Params{
		"photoset_id": id,
	})

	if err != nil {
		return nil, err
	}

	raw := &struct {
		PhotoSet struct {
			Id      string        `json:"id"`
			Primary string        `json:"primary"`
			Owner   string        `json:"owner"`
			Photos  []*AlbumPhoto `json:"photo"`
		} `json:"photoset"`
	}{}

	if err := Parse(response, raw); err != nil {
		return nil, err
	}

	response, err = client.Request("photosets.getInfo", Params{
		"photoset_id": id,
	})

	if err != nil {
		return nil, err
	}

	info := &struct {
		PhotoSet struct {
			Title       struct{
				Content string `json:"_content"`
			} `json:"title"`
			Description struct{
				Content string `json:"_content"`
			} `json:"description"`
		} `json:"photoset"`
	}{}

	if err := Parse(response, info); err != nil {
		return nil, err
	}

	return &Album{
		Id:      raw.PhotoSet.Id,
		Title:   info.PhotoSet.Title.Content,
		Desc:    info.PhotoSet.Description.Content,
		Primary: raw.PhotoSet.Primary,
		Owner:   raw.PhotoSet.Owner,
		Photos:  raw.PhotoSet.Photos,
	}, nil
}

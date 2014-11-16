package flickr

import (
	"strconv"
	"fmt"
)

type Photo struct {
	Id int
	Title string
	Description string
	PostTS int
	URL string
	Format string
	Views int

	Secret string
	Server int
	Farm int

	Username string
	UserFullName string
	UserId string
	UserLocation string
	UserIcon string
}

type PhotoRaw struct {
	Photo struct {
		Id string
		Secret string
		Server string
		Farm int
		DateUploaded string

		IsFavorite interface{}
		License interface{}
		Safety_Level interface{}
		Rotation interface{}

		OriginalSecret string
		OriginalFormat string

		Owner struct {
			NSID string
			Username string
			Realname string
			Location string
			IconServer string
			IconFarm int
			Path_Alias string
		}

		Title map[string]string
		Description map[string]string

		// who cares about this shit
		Visibility interface{}
		Dates interface{}
		Editability interface{}
		PublicEditability interface{}
		Usage interface{}
		Comments interface{}
		Notes interface{}
		People interface{}
		Tags interface{}

		URLs map[string][]map[string]string

		Media interface{}
		Stat interface{}
	}
}

func (client *Client) GetPhoto (id int) (*Photo, error) {
	response, err := client.Request("photos.getInfo", Params{
		"photo_id": fmt.Sprintf("%d", id),
	})

	if err != nil {
		return nil, err
	}

	raw := &PhotoRaw{}
	err = Parse(response, raw)

	if err != nil {
		return nil, err
	}

	date, err := strconv.Atoi(raw.Photo.DateUploaded)

	if err != nil {
		return nil, err
	}

	server, err := strconv.Atoi(raw.Photo.Server)

	if err != nil {
		return nil, err
	}

	icon := fmt.Sprintf("https://farm%d.staticflickr.com/%s/buddyicons/%s_r.jpg", raw.Photo.Owner.IconFarm, raw.Photo.Owner.IconServer, raw.Photo.Owner.NSID)

	return &Photo{
		Id: id,
		Title: raw.Photo.Title["_content"],
		Description: raw.Photo.Description["_content"],
		PostTS: date,
		Format: raw.Photo.OriginalFormat,
		Secret: raw.Photo.OriginalSecret,
		Server: server,
		Farm: raw.Photo.Farm,

		Username: raw.Photo.Owner.Path_Alias,
		UserFullName: raw.Photo.Owner.Realname,
		UserId: raw.Photo.Owner.NSID,
		UserLocation: raw.Photo.Owner.Location,
		UserIcon: icon,
	}, nil
}

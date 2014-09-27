package flickr

import (
	"fmt"
)

type User struct {
	Id string
	Name string
	Title string
	IconFarm int
	IconServer string
}

type UserRaw struct {
	User struct {
		Id string
		Username map[string]string
		Stat string
	}
}

func (client *Client) FindUser (name string) (*User, error) {
	response, err := client.Request("urls.lookupUser", Params{
		"url": fmt.Sprintf("http://flickr.com/photos/%s", name),
	})

	if err != nil {
		return nil, err
	}

	raw := &UserRaw{}
	err = Parse(response, raw)

	if err != nil {
		return nil, err
	}

	return &User{
		Id: raw.User.Id,
		Name: name,
	}, nil
}

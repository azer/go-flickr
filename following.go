package flickr

type FollowingRaw struct {
	Contacts struct {
		Contact []struct {
			IconFarm int
			IconServer string
			Ignored int
			NSID string
			RevIgnored int
			Username string
		}
	}
}

func (client *Client) Following (userId string) ([]User, error) {
	response, err := client.Request("contacts.getPublicList", Params{
		"user_id": userId,
	})

	raw := &FollowingRaw{}
	err = Parse(response, raw)

	if err != nil {
		return nil, err
	}

	following := []User{}

	for _, user := range raw.Contacts.Contact {
		following = append(following, User{
			Id: user.NSID,
			Title: user.Username,
			IconFarm: user.IconFarm,
			IconServer: user.IconServer,
		})
	}

	return following, nil
}

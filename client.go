package flickr

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

type Params map[string]string

type Client struct {
	Key   string
	Token string
	Sig   string
}

func (client *Client) Request (method string, params Params) ([]byte, error) {
	url := fmt.Sprintf("https://api.flickr.com/services/rest/?method=flickr.%s&api_key=%s&format=json&nojsoncallback=1", method, client.Key)

	if len(client.Token) > 0 {
		url = fmt.Sprintf("%s&auth_token=%s", url, client.Token)
	}

	if len(client.Sig) > 0 {
		url = fmt.Sprintf("%s&auth_sig=%s", url, client.Sig)
	}

	for key, value := range params {
		url = fmt.Sprintf("%s&%s=%s", url, key, value)
	}

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

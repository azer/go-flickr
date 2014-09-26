package flickr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response map[string]interface{}
type Params map[string]string
type Request func(method string, params *Params) (Response, error)

type Options struct {
	Key string
	Token string
	Sig string
}

func Client (options *Options) Request {
	return func (method string, params *Params) (Response, error) {
		url := fmt.Sprintf("flickr.%s&api_key=%s&format=json&nojsoncallback=1", method, options.Key)
		url = fmt.Sprintf("https://api.flickr.com/services/rest/?method=%s", url)

		if len(options.Token) > 0 {
			url = fmt.Sprintf("%s&auth_token=%s", url, options.Token)
		}

		if len(options.Sig) > 0 {
			url = fmt.Sprintf("%s&auth_sig=%s", url, options.Sig)
		}

		response, err := http.Get(url)
    defer response.Body.Close()

		if err != nil {
			return nil, err
		}

		data := Response{}
		dec := json.NewDecoder(response.Body)
		dec.Decode(&data)

		return data, nil
	}
}

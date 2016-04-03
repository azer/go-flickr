package flickr

import (
	"fmt"
)

var URL_KEYS = map[string]string{
	"square":      "s",
	"largeSquare": "q",
	"thumbnail":   "t",
	"small":       "m",
	"small320":    "n",
	"medium":      "",
	"medium640":   "z",
	"medium800":   "c",
	"large":       "b",
	"large1600":   "h",
	"large2048":   "k",
}

func GenerateURLs(id string, farm int, secret, server, format string) map[string]string {
	result := map[string]string{}
	for name, letter := range URL_KEYS {
		result[name] = GenerateURL(letter, id, farm, secret, server, format)
	}
	return result
}

func GenerateURL(key, id string, farm int, secret, server, format string) string {
	return fmt.Sprintf("https://farm%d.staticflickr.com/%s/%s_%s_%s.%s", farm, server, id, secret, key, format)
}

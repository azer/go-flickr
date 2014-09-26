## go-flickr-client

A minimalistic Flickr API client for Go

```go
import (
  "github.com/azer/go-flickr-client"
)

client := flickr.Client(&flickr.Options{
  Key: "key",
  Token: "token",
  Sig: "sig",
})

response, err := client("people.findByUsername", &flickr.Params{ "username": "azer" })

response["user"]["id"]
// => 13517180@N00
```

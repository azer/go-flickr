## go-flickr-client

A minimalistic client for Go

```go
import (
  "github.com/azer/go-flickr-client"
)

client := flickr.Client(&flickr.Options{
  key: "key",
  token: "token",
  sig: "sig"
})

response, err := client("people.findByUsername", { username: "azer" })

response["user"]["id"]
// => 13517180@N00
```

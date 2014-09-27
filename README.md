## go-flickr-client

A minimalistic Flickr API client for Go

## Install

```
$ go get github.com/azer/go-flickr
```

## API Reference

## Request

```go
import (
  "github.com/azer/go-flickr"
)

client := &flickr.Client{
  Key: "key",
  Token: "token", // optional
  Sig: "sig", // optional
}

response, err := client.Get("people.findByUsername", &flickr.Params{ "username": "azer" })
// => {"user":{"id":"98269877@N00", "nsid":"98269877@N00", "username":{"_content":"azerbike"}}, "stat":"ok"}
```

## FindUser

Find user by name.

```go
user, err := client.FindUser("azer")

user.Id
// => 123124324

user.Name
// => azer
```

## Favs

List given user's favorites

```go
userId := 123123123

favs, err := client.Favs(userId)
```

## Following

List the people given user follows on Flickr

```go
userId := 123123123

following, err := client.Following(userId)
```

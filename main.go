package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/r3labs/sse/v2"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "15879e0a-1bd0-43fa-a1c5-4314e868db57",
		ClientSecret: "YFzpmts8grh5caFmj4W0e80zE8j3DpkR",
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.mcs3.miele.com/thirdparty/login",
			TokenURL: "https://api.mcs3.miele.com/thirdparty/token",
		},
	}

	client := sse.NewClient("https://api.mcs3.miele.com/v1/devices/all/events")
	// client.Headers["Accept-Language"] = "en"

	token := &oauth2.Token{
		AccessToken:  "GB_7c3258712ec54ea1e0cceca8dc082298",
		RefreshToken: "GB_7e49d2cb0a9634e1b745e599d9fd3ea2",
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(-time.Minute),
	}
	ts := conf.TokenSource(ctx, token)
	token, err := ts.Token()
	if err != nil {
		log.Fatalln(err)
	}

	client.Connection = conf.Client(ctx, token)

	client.Subscribe("", func(msg *sse.Event) {
		// Got some data!
		fmt.Println(string(msg.Data))
	})
}

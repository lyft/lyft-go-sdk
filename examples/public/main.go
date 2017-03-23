package main

import (
	"context"

	"github.com/lyft/lyft-go-sdk/lyft"

	"log"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	config := clientcredentials.Config{
		ClientID:     "YOUR-CLIENT-ID",
		ClientSecret: "YOUR-CLIENT-SECRET",
		TokenURL:     "https://api.lyft.com/oauth/token",
		Scopes:       []string{"public"},
	}

	httpClient := config.Client(context.Background())
	client := lyft.NewAPIClient(httpClient, "sample-app")

	opt := map[string]interface{}{
		"endLat":   37.7884,
		"endLng":   -122.4076,
		"rideType": string(lyft.RideTypeLyft),
	}
	result, _, err := client.PublicApi.GetCost(37.7763, -122.3918, opt)

	if err != nil {
		log.Fatal(err)
	} else {
		spew.Dump(result)
	}
}

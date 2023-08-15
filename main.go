package main

import (
	"context"
	"go.einride.tech/here/geocodingsearchv7"
	"log"
	"net/http"
)

const apiKey = "API_KEY"

func main() {
	ctx := context.Background()
	// Create an authenticated client
	geocodingClient := geocodingsearchv7.NewClient(
		geocodingsearchv7.NewAPIKeyHTTPClient(apiKey, http.DefaultClient.Transport),
	)

	response, err := geocodingClient.ReverseGeocoding.ReverseGeocoding(
		ctx,
		&geocodingsearchv7.ReverseGeocodingRequest{
			GeoPosition: &geocodingsearchv7.GeoWaypoint{
				Lat:  39.469908,
				Long: -0.376288,
			},
		},
	)

	if err != nil {
		panic(err)
	}

	for _, item := range response.Items {
		log.Println(item.Address.Label)
	}
}

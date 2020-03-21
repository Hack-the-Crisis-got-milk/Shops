package utils

import (
	"github.com/Hack-the-Crisis-got-milk/Shops/environment"
	"googlemaps.github.io/maps"
)

func NewGoogleMapsClient(env *environment.Env) (*maps.Client, error) {
	return maps.NewClient(maps.WithAPIKey(env.Get(environment.GooglePlacesAPIKey)))
}

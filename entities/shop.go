package entities

import (
	"googlemaps.github.io/maps"
)

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Shop struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Loc      maps.LatLng `json:"loc"`
	Address  string      `json:"address"`
	OpenNow  bool        `json:"open_now"`
	Photo    string      `json:"photo"`
	Distance uint        `json:"distance"`
}

func ConvertPlacesSearchResponseToShops(response maps.PlacesSearchResponse, startLocation maps.LatLng) []Shop {
	shops := []Shop{}
	for _, response := range response.Results {
		shop := Shop{
			ID:       response.PlaceID,
			Name:     response.Name,
			Loc:      response.Geometry.Location,
			Address:  response.FormattedAddress,
			Distance: 0,
		}

		if len(response.Photos) > 0 {
			shop.Photo = response.Photos[0].PhotoReference
		}

		if response.OpeningHours != nil {
			shop.OpenNow = *response.OpeningHours.OpenNow
		}

		shops = append(shops, shop)
	}

	return shops
}

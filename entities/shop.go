package entities

import (
	"context"
	"fmt"
	"googlemaps.github.io/maps"
	"image/jpeg"
	"os"
)

const IMAGES_PATH = "./static/images/"

type Shop struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Loc      maps.LatLng `json:"loc"`
	Address  string      `json:"address"`
	OpenNow  bool        `json:"open_now"`
	Photo    string      `json:"photo"`
	Distance uint        `json:"distance"`
}

func ConvertPlacesSearchResponseToShops(ctx context.Context, response maps.PlacesSearchResponse, startLocation maps.LatLng, gClient *maps.Client) []Shop {
	shops := []Shop{}
	for _, response := range response.Results {
		photo, _ := getShopPhoto(ctx, response.PlaceID, gClient, response.Photos)

		shop := Shop{
			ID:       response.PlaceID,
			Name:     response.Name,
			Loc:      response.Geometry.Location,
			Address:  response.FormattedAddress,
			Photo:    photo,
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

func getShopPhoto(ctx context.Context, shopId string, gClient *maps.Client, photos []maps.Photo) (string, error) {
	os.MkdirAll(IMAGES_PATH, os.ModePerm)
	_, err := os.Open(IMAGES_PATH + fmt.Sprintf("%s.jpg", shopId))
	if err == nil {
		fmt.Println("picture found")
		return IMAGES_PATH + fmt.Sprintf("%s.jpg", shopId), nil
	}

	if len(photos) == 0 {
		return "", nil
	}

	res, err := gClient.PlacePhoto(ctx, &maps.PlacePhotoRequest{
		PhotoReference: photos[0].PhotoReference,
		MaxHeight:      800,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	img, err := res.Image()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	f, err := os.Create("./static/images/" + shopId + ".jpg")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()
	jpeg.Encode(f, img, nil)

	return IMAGES_PATH + fmt.Sprintf("%s.jpg", shopId), nil
}

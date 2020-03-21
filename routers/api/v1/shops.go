package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hack-the-Crisis-got-milk/Shops/entities"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"googlemaps.github.io/maps"
)

const LNG_KEY = "lng"
const LAT_KEY = "lat"
const FILTERS_KEY = "filters"
const RADIUS_KEY = "radius"

const GROCERY_STORE_SEARCH_KEYWORD = "grocery store"
const PHARMACY_SEARCH_KEYWORD = "pharmacy"

const DEFAULT_RADIUS = 1000

func newGetNearbyShopsRequest(ctx *gin.Context) (getNearbyShopsRequest, error) {
	if ctx.Query(LAT_KEY) == "" || ctx.Query(LNG_KEY) == "" {
		return getNearbyShopsRequest{}, errors.New(fmt.Sprintf("both %s and %s must be provided", LNG_KEY, LAT_KEY))
	}

	request := getNearbyShopsRequest{
		Loc:     maps.LatLng{},
		Filters: []entities.Filter{},
	}
	var err error
	request.Loc.Lat, err = strconv.ParseFloat(ctx.Query(LAT_KEY), 64)
	if err != nil {
		return getNearbyShopsRequest{}, errors.New(fmt.Sprintf("could not convert %s to %T", LAT_KEY, request.Loc.Lat))
	}

	request.Loc.Lng, err = strconv.ParseFloat(ctx.Query(LNG_KEY), 64)
	if err != nil {
		return getNearbyShopsRequest{}, errors.New(fmt.Sprintf("could not convert %s to %T", LNG_KEY, request.Loc.Lng))
	}

	if ctx.Query(RADIUS_KEY) == "" {
		request.Radius = DEFAULT_RADIUS
	} else {
		request.Radius, err = strconv.ParseUint(ctx.Query(RADIUS_KEY), 10, 64)
		if err != nil {
			return getNearbyShopsRequest{}, errors.New(fmt.Sprintf("could not convert %s to %T", RADIUS_KEY, request.Radius))
		}
	}

	err = json.Unmarshal([]byte(ctx.Query(FILTERS_KEY)), &request.Filters)
	if err != nil {
		return getNearbyShopsRequest{}, errors.New(fmt.Sprintf("could not convert %s to %T", FILTERS_KEY, request.Filters))
	}

	return request, nil
}

func (r *apiV1Router) getShopsWithinRadius(ctx *gin.Context, startpoint maps.LatLng, radius uint, searchKeyword string, placeType maps.PlaceType) ([]entities.Shop, error) {
	response, err := r.gClient.NearbySearch(ctx, &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: startpoint.Lat,
			Lng: startpoint.Lng,
		},
		Radius:  radius,
		Keyword: searchKeyword,
		Type:    placeType,
	})

	r.logger.Info("response", zap.Any("shops", response))
	if err != nil {
		return nil, err
	}

	return entities.ConvertPlacesSearchResponseToShops(response, startpoint), nil
}

func (r *apiV1Router) GetNearbyShops(ctx *gin.Context) {
	request, err := newGetNearbyShopsRequest(ctx)
	if err != nil {
		r.logger.Debug("could not parse GetNearbyShops request", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, newErrorResponse("could not parse request params", err))
		return
	}

	shops, err := r.getShopsWithinRadius(ctx, request.Loc, uint(request.Radius), GROCERY_STORE_SEARCH_KEYWORD, "grocery_or_supermarket")
	if err != nil {
		r.logger.Error("could not fetch grocery stores", zap.Error(err))
	}
	pharmacies, err := r.getShopsWithinRadius(ctx, request.Loc, uint(request.Radius), PHARMACY_SEARCH_KEYWORD, maps.PlaceTypePharmacy)
	if err != nil {
		r.logger.Error("could not fetch pharmacies", zap.Error(err))
	}

	shops = append(shops, pharmacies...)

	ctx.JSON(http.StatusOK, getNearbyShopsResponse{
		Shops:    shops,
	})
}

func (r *apiV1Router) GetAllShops(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, getNearbyShopsResponse{
		Shops: []entities.Shop{
			{
				Name:    "Rimi",
				Address: "Europos pr. 43, Kaunas 46329, Lituania",
				Loc: maps.LatLng{
					Lat: 54.8759003,
					Lng: 23.9120662,
				},
				OpenNow: true,
				Photo:   "https://lh3.googleusercontent.com/p/AF1QipMkw2wH2iv11UZBWBW0L2Ki5Ei7cguRdhXWjjvl=s1600-w400",
			},
		},
	})
}

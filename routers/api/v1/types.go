package v1

import (
	"fmt"
	"github.com/Hack-the-Crisis-got-milk/Shops/entities"
	"googlemaps.github.io/maps"
)

// Response is the basic model for an API response
type Response struct {
	Error string `json:"error,omitempty"`
}

func newErrorResponse(message string, err error) Response {
	return Response{Error: fmt.Sprintf("%s: %s", message, err.Error())}
}

type getItemGroupsResponse struct {
	Response
	ItemGroups []entities.ItemGroup `json:"item_groups"`
}

type getNearbyShopsRequest struct {
	Loc     maps.LatLng       `json:"loc"`
	Radius  uint64            `json:"radius"`
	Filters []entities.Filter `json:"filters"`
}

type getNearbyShopsResponse struct {
	Response
	Shops []entities.Shop `json:"shops"`
}

type getAllShopsResponse struct {
	Response
	Shops []entities.Shop `json:"shops"`
}

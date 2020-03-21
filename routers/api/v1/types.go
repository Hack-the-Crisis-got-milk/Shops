package v1

import "github.com/Hack-the-Crisis-got-milk/Shops/entities"

// Response is the basic model for an API response
type Response struct {
	Error  string `json:"error,omitempty"`
}

type getItemGroupsResponse struct {
	Response
	ItemGroups []entities.ItemGroup `json:"item_groups"`
}

type getNearbyShopsResponse struct {
	Response
	Shops []entities.Shop `json:"shops"`
}

type getAllShopsResponse struct {
	Response
	Shops []entities.Shop `json:"shops"`
}


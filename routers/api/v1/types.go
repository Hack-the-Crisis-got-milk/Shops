package v1

import "github.com/Hack-the-Crisis-got-milk/Shops/entities"

// Response is the basic model for an API response
type Response struct {
	Error  string `json:"error"`
}

type getItemGroupsResponse struct {
	Response
	ItemGroups []entities.ItemGroup `json:"item_groups"`
}


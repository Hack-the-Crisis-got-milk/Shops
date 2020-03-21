package v1

import (
	"github.com/Hack-the-Crisis-got-milk/Shops/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *apiV1Router) GetNearbyShops(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, getNearbyShopsResponse{
		Shops:    []entities.Shop{
			entities.Shop{
				Name: "Rimi",
				Address: "Europos pr. 43, Kaunas 46329, Lituania",
				Loc: entities.Location{
					Lat: 54.8759003,
					Long: 23.9120662,
				},
				OpenNow: true,
				Photo: "https://lh3.googleusercontent.com/p/AF1QipMkw2wH2iv11UZBWBW0L2Ki5Ei7cguRdhXWjjvl=s1600-w400",
			},
		},
	})
}

func (r *apiV1Router) GetAllShops(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, getNearbyShopsResponse{
		Shops:    []entities.Shop{
			entities.Shop{
				Name: "Rimi",
				Address: "Europos pr. 43, Kaunas 46329, Lituania",
				Loc: entities.Location{
					Lat: 54.8759003,
					Long: 23.9120662,
				},
				OpenNow: true,
				Photo: "https://lh3.googleusercontent.com/p/AF1QipMkw2wH2iv11UZBWBW0L2Ki5Ei7cguRdhXWjjvl=s1600-w400",
			},
		},
	})
}


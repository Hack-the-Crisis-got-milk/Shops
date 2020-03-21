package v1

import (
	"github.com/Hack-the-Crisis-got-milk/Shops/config"
	"github.com/Hack-the-Crisis-got-milk/Shops/environment"
	"github.com/Hack-the-Crisis-got-milk/Shops/routers/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"googlemaps.github.io/maps"
)

// APIV1Router is the router for v1 of the API
type APIV1Router interface {
	models.Router
	GetNearbyShops(*gin.Context)
	GetAllShops(*gin.Context)
	GetItemGroups(*gin.Context)
}

type apiV1Router struct {
	models.BaseRouter
	logger  *zap.Logger
	env     *environment.Env
	cfg     *config.AppConfig
	gClient *maps.Client
}

// NewAPIV1Router creates a APIV1Router
func NewAPIV1Router(logger *zap.Logger, env *environment.Env, cfg *config.AppConfig, gClient *maps.Client) APIV1Router {
	return &apiV1Router{
		logger:  logger,
		env:     env,
		cfg:     cfg,
		gClient: gClient,
	}
}

// RegisterRoutes registers all of the API's (v1) routes to the given router group
func (r apiV1Router) RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", r.Heartbeat)

	shopsGroup := routerGroup.Group("/shops")
	shopsGroup.GET("/", r.GetAllShops)
	shopsGroup.GET("/nearby", r.GetNearbyShops)

	itemGroupsGroup := routerGroup.Group("/itemgroups")
	itemGroupsGroup.GET("/", r.GetItemGroups)
}

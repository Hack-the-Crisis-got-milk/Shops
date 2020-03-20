package v1

import (
	"github.com/Hack-the-Crisis-got-milk/Shops/repositories"
	"github.com/gin-gonic/gin"
	"github.com/Hack-the-Crisis-got-milk/Shops/environment"
	"github.com/Hack-the-Crisis-got-milk/Shops/routers/models"
	"go.uber.org/zap"
)

// APIV1Router is the router for v1 of the API
type APIV1Router interface {
	models.Router
	GetNearbyShops(*gin.Context)
	GetAllShops(*gin.Context)
}

type apiV1Router struct {
	models.BaseRouter
	logger       *zap.Logger
	env          *environment.Env
	itemGroupRepo *repositories.ItemGroupRepository
}

// NewAPIV1Router creates a APIV1Router
func NewAPIV1Router(logger *zap.Logger, env *environment.Env, itemGroupRepo *repositories.ItemGroupRepository) APIV1Router {
	return &apiV1Router{
		logger:       logger,
		env:          env,
		itemGroupRepo: itemGroupRepo,
	}
}

// RegisterRoutes registers all of the API's (v1) routes to the given router group
func (r apiV1Router) RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", r.Heartbeat)
	routerGroup.GET("/shops", r.GetAllShops)
	routerGroup.GET("/shops/nearby", r.GetNearbyShops)
}

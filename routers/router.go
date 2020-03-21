package routers

import (
	v1 "github.com/Hack-the-Crisis-got-milk/Shops/routers/api/v1"
	"github.com/Hack-the-Crisis-got-milk/Shops/routers/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// MainRouter is router to connect all routers used by the app
type MainRouter interface {
	models.Router
}

type mainRouter struct {
	models.BaseRouter
	logger *zap.Logger
	apiV1  v1.APIV1Router
}

// NewMainRouter creates a new MainRouter
func NewMainRouter(logger *zap.Logger, apiV1 v1.APIV1Router) MainRouter {
	return &mainRouter{
		logger: logger,
		apiV1:  apiV1,
	}
}

// RegisterRoutes registers all of the app's routes
func (r *mainRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	apiV1Group := routerGroup.Group("/api/v1")
	r.apiV1.RegisterRoutes(apiV1Group)
}

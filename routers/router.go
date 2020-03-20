package routers

import (
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
}

// NewMainRouter creates a new MainRouter
func NewMainRouter(logger *zap.Logger) MainRouter {
	return &mainRouter{
		logger: logger,
	}
}

// RegisterRoutes registers all of the app's routes
func (r *mainRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
}

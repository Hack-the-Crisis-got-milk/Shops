package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *apiV1Router) GetItemGroups(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "Not implemented!")
}
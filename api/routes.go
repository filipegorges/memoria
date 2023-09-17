package api

import (
	"github.com/filipegorges/memoria/api/history"
	"github.com/gin-gonic/gin"
)

func BuildRoutes(routes *gin.RouterGroup) {
	v1 := routes.Group("/v1")

	history.Handlers(v1)
}

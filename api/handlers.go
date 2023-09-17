package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) Handlers(routes *gin.RouterGroup) {
	h := routes.Group("/history")
	h.GET("/", a.findAllHistory)
}

func (a *API) findAllHistory(c *gin.Context) {
	a.logger.Info("findAllHistory")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

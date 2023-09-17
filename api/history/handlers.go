package history

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handlers(routes *gin.RouterGroup) {
	h := routes.Group("/history")
	h.GET("/", findAllHistory)
}

func findAllHistory(c *gin.Context) {
	log.Println("findAllHistory")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

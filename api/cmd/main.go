package main

import (
	"github.com/filipegorges/memoria/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	api.BuildRoutes(&r.RouterGroup)

	if err := r.Run(); err != nil {
		panic(err)
	}
}

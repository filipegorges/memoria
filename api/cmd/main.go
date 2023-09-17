package main

import (
	"github.com/filipegorges/memoria/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.New()
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	api := api.NewAPI(r, logger)
	api.BuildRoutes()
	api.BuildMiddlewares()

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type API struct {
	logger *zap.Logger
	srv    *gin.Engine
}

func NewAPI() *API {
	r := gin.New()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	a := &API{srv: r, logger: logger}
	r.Use(a.zapLogging())
	return a
}

func (a *API) Run(addr ...string) error {
	a.BuildRoutes()
	a.srv.Use(a.zapLogging())

	if err := a.srv.Run(":8080"); err != nil {
		return err
	}

	return nil
}

func (a *API) Close() {
	a.logger.Sync()
}

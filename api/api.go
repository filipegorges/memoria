package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type API struct {
	logger *zap.Logger
	srv    *gin.Engine
}

func NewAPI(srv *gin.Engine, logger *zap.Logger) *API {
	return &API{
		srv:    srv,
		logger: logger,
	}
}

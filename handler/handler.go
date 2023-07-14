package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/service"
)

func CreateRouter(svc *service.ServiceRegistry) *gin.Engine {
	core.Log.Debug("Creating Router...")
	router := gin.Default()

	// setup each handler by adding it to the router
	setupLeague(router, svc.League)

	return router
}
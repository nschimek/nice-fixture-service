package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/service"
)

type league struct {
	gr *gin.Engine
	svc service.League
}


func NewLeague(gr *gin.Engine, svc service.League) {
	
}
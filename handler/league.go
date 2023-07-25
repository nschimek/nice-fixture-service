package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/service"
)

const leagueEndpoint = "/leagues"

type League interface {
	GetByParams(c *gin.Context)
	GetById(c *gin.Context)
}

type league struct {
	svc service.League
}

func setupLeague(rg *gin.RouterGroup, svc service.League) {
	core.Log.Debug("Setting up League Handler...")

	h := &league{svc: svc}
	g := rg.Group(leagueEndpoint)
	g.GET("", h.GetByParams)
	g.GET("/:id", h.GetById)
}

func (h *league) GetByParams(c *gin.Context) {
	p := model.LeagueParams{}
	if ok := bind(c.JSON, c.ShouldBind, &p); !ok {
		return 
	}
	r, err := h.svc.GetByParams(p)
	jsonResult[[]model.League](c.JSON, &r, err)
}

func (h *league) GetById(c *gin.Context) {
	p := idParam{}
	if ok := bind(c.JSON, c.ShouldBindUri, &p); !ok {
		return
	}
	r, err := h.svc.GetById(p.ID)
	jsonResult[model.League](c.JSON, r, err)
}
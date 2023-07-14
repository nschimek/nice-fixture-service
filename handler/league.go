package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/service"
)

const leagueEndpoint = "leagues"

type League interface {
	GetByParams(c *gin.Context)
	GetById(id int)
}

type league struct {
	gr *gin.Engine
	svc service.League
}

func setupLeague(gr *gin.Engine, svc service.League) {
	bp := fmt.Sprintf("%s/%s", core.ApiBasePath, leagueEndpoint)

	core.Log.WithField("path", bp).Debug("Setting up League Handler...")

	h := &league{gr: gr, svc: svc}
	g := h.gr.Group(bp)
	g.GET("", h.GetByParams)
	g.GET("/:id", h.GetById)
}

func (h *league) GetByParams(c *gin.Context) {
	p := model.LeagueParams{}
	if ok := bind(c, c.ShouldBind, &p); !ok {
		return 
	}
	r, err := h.svc.GetByParams(p)
	jsonResult[[]model.League](c, &r, err)
}

func (h *league) GetById(c *gin.Context) {
	p := idParam{}
	if ok := bind(c, c.ShouldBindUri, &p); !ok {
		return
	}
	r, err := h.svc.GetById(p.ID)
	jsonResult[model.League](c, r, err)
}
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/service"
)

const teamEndpoint = "/teams"

type Team interface {
	GetByParams(c *gin.Context)
	GetById(c *gin.Context)
}

type team struct {
	svc service.Team
}

func setupTeam(rg *gin.RouterGroup, svc service.Team) {
	core.Log.Debug("Setting up Team Handler...")

	h := &team{svc: svc}
	g := rg.Group(teamEndpoint)

	g.GET("", h.GetByParams)
	g.GET(idBindPath, h.GetById)
}

func (h *team) GetByParams(c *gin.Context) {
	p := model.TeamParams{}
	if ok := bind(c.JSON, c.ShouldBind, &p); !ok {
		return
	}
	r, err := h.svc.GetByParams(p)
	jsonResult[[]model.Team](c.JSON, &r, err)
}

func (h *team) GetById(c *gin.Context) {
	p := idParam{}
	if ok := bind(c.JSON, c.ShouldBindUri, &p); !ok {
		return
	}
	r, err := h.svc.GetById(p.ID)
	jsonResult[model.Team](c.JSON, r, err)
}
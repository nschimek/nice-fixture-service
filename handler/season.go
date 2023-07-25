package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/service"
)

const seasonEndpoint = "/seasons"

type Season interface {
	GetAll(c *gin.Context)
}

type season struct {
	svc service.Season
}

func setupSeason(rg *gin.RouterGroup, svc service.Season) {
	core.Log.Debug("Setting up Season Handler...")

	h := &season{svc: svc}
	g := rg.Group(seasonEndpoint)
	g.GET("", h.GetAll)
}

func (h *season) GetAll(c *gin.Context) {
	r, err := h.svc.GetAll()
	jsonResult[[]model.Season](c.JSON, &r, err)
}

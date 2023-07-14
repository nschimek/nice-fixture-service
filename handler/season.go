package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/service"
)

const seasonEndpoint = "seasons"

type Season interface {
	GetByParams(c *gin.Context)
	GetById(id int)
}

type season struct {
	gr *gin.Engine
	svc service.Season
}

func setupSeason(gr *gin.Engine, svc service.Season) {
	bp := fmt.Sprintf("%s/%s", core.ApiBasePath, seasonEndpoint)

	core.Log.WithField("path", bp).Debug("Setting up Season Handler...")

	h := &season{gr: gr, svc: svc}
	g := h.gr.Group(bp)
	g.GET("", h.GetAll)
}

func (h *season) GetAll(c *gin.Context) {
	r, err := h.svc.GetAll()
	jsonResult[[]model.Season](c, &r, err)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/service"
)

const leagueEndpoint = "leagues"

type League interface {
	GetByParams(c *gin.Context)
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
}

func (h *league) GetByParams(c *gin.Context) {
	p := model.LeagueParams{}
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if r, err := h.svc.GetByParams(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, r)
	}
}
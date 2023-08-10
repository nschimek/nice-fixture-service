package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/service"
)

const fixtureEndpoint = "/fixtures"

type Fixture interface {
	GetByParams(c *gin.Context)
}

type fixture struct {
	svc service.Fixture
}

func setupFixture(rg *gin.RouterGroup, svc service.Fixture) {
	core.Log.Debug("Setting up Fixture Handler...")

	h := &fixture{svc: svc}
	g := rg.Group(fixtureEndpoint)
	g.GET("", h.GetByParams)
}

func (h *fixture) GetByParams(c *gin.Context) {
	r, err := h.svc.GetAll()
	jsonResult[[]model.Fixture](c.JSON, &r, err)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/service"
)

type idParam struct {
	ID int `uri:"id" binding:"required,gt=0"` 
}

func CreateRouter(svc *service.ServiceRegistry) *gin.Engine {
	core.Log.Debug("Creating Router...")
	router := gin.Default()

	// setup each handler by adding it to the router
	setupLeague(router, svc.League)

	return router
}

// Handler helper functions

// Given a Gin context, a result, and a potential error, add either of the following to the Gin context JSON response:
//  - the result in JSON format along with a status 200
//  - the error in JSON format as an internal server error w/status code 500
func jsonResult[T any](c *gin.Context, result *T, err error) {
	if err != nil {
		errorResult(c, rest_error.New(rest_error.Internal))
	} else if result == nil {
		errorResult(c, rest_error.New(rest_error.NotFound))
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// Convert the rest error into a Gin JSON result
func errorResult(c *gin.Context, err *rest_error.Error) {
	c.JSON(err.Code, err)
}

// Perform binding validation on the given params using the given binding function.
// Returns true if validation passes, and false if it does not.  
// Also adds the error to the Gin context JSON response if validation fails.
func bind(c *gin.Context, bindingFunc func(obj any) error, params any) bool {
	// run the binding func against the incoming params and check for validation errors
	if err := bindingFunc(params); err != nil {
		core.Log.Warnf("Binding valdation failed: %v", err)
		if errs, ok := err.(validator.ValidationErrors); ok {
			details := []string{}
			for _, err := range errs {
				details = append(details, fmt.Sprintf("%s [%s]: tag %s failed validation", err.Field(), err.Param(), err.Tag()))
			}
			errorResult(c, rest_error.NewWithDetails(rest_error.BadRequest, details))
			return false
		}
		errorResult(c, rest_error.NewWithDetail(rest_error.BadRequest, err.Error()))
		return false
	}
	return true
}
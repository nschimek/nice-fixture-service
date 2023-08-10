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

const idBindPath = "/:id"

type idParam struct {
	ID int `uri:"id" binding:"required,gt=0"` 
}

func CreateRouter(svc *service.ServiceRegistry) *gin.Engine {
	core.Log.Debug("Creating Router...")
	router := gin.Default()

	v1 := router.Group(core.ApiBasePath(1))

	// setup each handler by adding it to the router group
	setupLeague(v1, svc.League)
	setupSeason(v1, svc.Season)
	setupTeam(v1, svc.Team)
	setupFixture(v1, svc.Fixture)

	return router
}

// Handler helper functions

// Using the given JSON response function (gin.Context.JSON), along with a result and a potential error, handle one of the following:
//  - the result in JSON format along with a status 200 if the result is not nil
//  - an error in JSON format along with a status of 404 if the result is nil
//  - the error in JSON format if its not nill
func jsonResult[T any](jsonResponseFunc func(code int, obj any), result *T, err *rest_error.Error) {
	if err != nil {
		errorResult(jsonResponseFunc, err)
	// the use of generics allows us to force a pointer, and now we can check for nil
	} else if result == nil {
		errorResult(jsonResponseFunc, rest_error.NewNotFound())
	} else {
		jsonResponseFunc(http.StatusOK, result)
	}
}

// Perform binding validation on the given params using the given binding function (from gin.Context)
// Returns true if validation passes, and false if it does not.  
// Also handles the error using the JSON func if the validation fails.
func bind(jsonResponseFunc func(code int, obj any), bindingFunc func(obj any) error, params any) bool {
	// run the binding func against the incoming params and check for validation errors
	if err := bindingFunc(params); err != nil {
		core.Log.Warnf("Binding valdation failed: %v", err)
		if errs, ok := err.(validator.ValidationErrors); ok {
			details := []string{}
			for _, err := range errs {
				details = append(details, fmt.Sprintf("%s [%s]: tag %s failed validation", err.Field(), err.Param(), err.Tag()))
			}
			errorResult(jsonResponseFunc, rest_error.NewBadRequest(details...))
			return false
		}
		errorResult(jsonResponseFunc, rest_error.NewBadRequest(err.Error()))
		return false
	}
	return true
}

// Handle the REST error with the JSON Handler function
func errorResult(jsonResponseFunc func(code int, obj any), err *rest_error.Error) {
	jsonResponseFunc(err.Code, err)
}
package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/services"
	"net/http"
)

type errorBody struct {
	Message string `json:"message"`
}

func (a *Api) errorHandler(err error, c echo.Context) {
	switch err.(type) {
	case *services.ValidationError:
		c.JSON(http.StatusUnprocessableEntity, &errorBody{err.Error()})
	case *services.EntityNotFoundError:
		c.JSON(http.StatusNotFound, &errorBody{err.Error()})
	case *services.InvalidQueryError:
		c.JSON(http.StatusBadRequest, &errorBody{err.Error()})
	case *services.UnprocessableBodyError:
		c.JSON(http.StatusBadRequest, &errorBody{err.Error()})
	}
}

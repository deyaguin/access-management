package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gitlab/nefco/access-management-system/src/services"
	"go.uber.org/zap"
	"net/http"
)

type errorBody struct {
	Message string `json:"message"`
}

func (a *API) errorHandler(
	err error,
	c echo.Context,
) {
	switch err.(type) {
	case *services.ValidationError:
		Log.Warn(
			"validation failed",
			zap.Error(err),
		)
		c.JSON(http.StatusUnprocessableEntity, &errorBody{err.Error()})
	case *InvalidQueryError:
		Log.Warn(
			"invalid query",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, &errorBody{err.Error()})
	case *UnprocessableBodyError:
		Log.Warn(
			"unprocessable body",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, &errorBody{err.Error()})
	case *services.EntityNotFoundError:
		Log.Error(
			"entity not found",
			zap.Error(err),
		)
		c.JSON(http.StatusNotFound, &errorBody{err.Error()})
	case *services.EntityCreateError:
		Log.Error(
			"entity create error",
			zap.Error(err),
		)
		c.NoContent(http.StatusInternalServerError)
	case *services.EntityUpdateError:
		Log.Error(
			"entity update error",
			zap.Error(err),
		)
		c.NoContent(http.StatusInternalServerError)
	case *services.EntityRemoveError:
		Log.Error(
			"entity remove error",
			zap.Error(err),
		)
		c.NoContent(http.StatusInternalServerError)
	case *services.GetEntitiesError:
		log.Error(
			"get entities error",
			zap.Error(err),
		)
		c.NoContent(http.StatusNotFound)
	}
}

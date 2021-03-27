package controllers

import (
	"context"
	"net/http"

	"github.com/Kourin1996/go-crud-api-sample/api/common/log"
	"github.com/Kourin1996/go-crud-api-sample/api/models/errors"
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

func getStatusAndMessage(err error, showsDetail bool) (int, string) {
	switch e := err.(type) {
	case errors.InvalidRequestError:
		return http.StatusBadRequest, e.Error()
	case errors.InvalidDataError:
		return http.StatusBadRequest, e.Error()
	case errors.NotFoundError:
		return http.StatusNotFound, e.Error()
	default:
		if showsDetail {
			return http.StatusInternalServerError, e.Error()
		}
		return http.StatusInternalServerError, "Internal Server Error"
	}
}

func logToSentry(err error, ctx echo.Context) {
	req := ctx.Request()
	hub := sentry.GetHubFromContext(req.Context())
	if hub == nil {
		hub = sentry.CurrentHub().Clone()
	}
	hub.Scope().SetRequest(req)
	hub.RecoverWithContext(
		context.WithValue(req.Context(), sentry.RequestContextKey, req),
		err,
	)
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewHTTPErrorHandler(cfg Config) func(err error, ctx echo.Context) {
	return func(err error, ctx echo.Context) {
		if err != nil {
			log.Error(err)
			logToSentry(err, ctx)
		}

		status, message := getStatusAndMessage(err, cfg.IsDebug)
		if !ctx.Response().Committed {
			if ctx.Request().Method == http.MethodHead {
				err = ctx.NoContent(status)
			} else {
				err = ctx.JSON(status, ErrorResponse{Status: status, Message: message})
			}
			if err != nil {
				log.Error(err)
			}
		}
	}
}

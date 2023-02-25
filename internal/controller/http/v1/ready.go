package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"softpro6/internal/usecase"
	"softpro6/pkg/logger"
)

type readyRoute struct {
	isAppReady usecase.IsAppReadyUseCase
	logger     logger.Interface
}

func newReadyRoute(router chi.Router, uc usecase.IsAppReadyUseCase, l logger.Interface) {
	rr := &readyRoute{
		isAppReady: uc,
		logger:     l,
	}

	router.Get("/ready", rr.ready)
}

// @Summary     Check readiness
// @Description 200 if ready, 500 if not
// @ID          ready
// @Tags  	    readiness
// @Accept      json
// @Produce     json
// @Success     200 {object} nil
// @Failure     500 {object} ErrResponse
// @Router      /ready [get]
func (rr *readyRoute) ready(w http.ResponseWriter, r *http.Request) {
	readiness, err := rr.isAppReady.Execute(r.Context())
	if err != nil {
		rr.logger.Error("http - ready - isAppReady.Execute", err)
		_ = render.Render(w, r, notReadyError(err))
		return
	}

	if readiness == nil {
		rr.logger.Info("http - ready - IsReady - readiness is nil")
		_ = render.Render(w, r, notReadyError(err))
		return
	}

	if !readiness.IsReady() {
		rr.logger.Info("http - ready - !IsReady", readiness.Reasons())
		_ = render.Render(w, r, notReadyError(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, nil)
}

func notReadyError(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "app is not ready",
		ErrorText:      "storage is not available or lines are not synced yet",
	}
}

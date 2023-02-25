package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"softpro6/internal/usecase"
	"softpro6/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Line Processor
// @description Checking app readiness
// @version     1.0
// @host        127.0.0.1:80
// @BasePath    /
func NewRouter(handler chi.Router, isAppReady usecase.IsAppReadyUseCase, l logger.Interface) {
	setMiddlewares(handler)
	setSwagger(handler)
	newReadyRoute(handler, isAppReady, l)
}

func setMiddlewares(router chi.Router) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
}

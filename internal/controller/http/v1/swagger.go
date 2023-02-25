package v1

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

func setSwagger(r chi.Router) {
	SwaggerFiles(r, "/swag")

	//swagger.SwaggerInfo.Host = host        //todo issue or make feature

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swag/swagger.json"), //The url pointing to API definition
	))
}

func SwaggerFiles(r chi.Router, path string) {
	_, b, _, _ := runtime.Caller(0)
	swaggerFolder := filepath.Join(b, "../../../../../api/swagger")

	if strings.ContainsAny(path, "{}*") {
		panic("SwaggerFiles does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rCtx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rCtx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(http.Dir(swaggerFolder)))
		fs.ServeHTTP(w, r)
	})
}

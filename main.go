package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/inquizarus/golam"
	"inquizarus.dev/site/config"
	"inquizarus.dev/site/routes"
)

func main() {
	site := config.Site{
		Name: "site",
		Options: config.Options{
			HTTPPort:    "80",
			StaticDir:   "statics",
			TemplateDir: "templates",
		},
	}
	config.Load(&site)
	cfg := golam.Config{
		UsePort: site.Options.HTTPPort,
		Routes: []golam.Route{
			routes.MakeRootRoute(site),
			routes.MakeJavaScriptRoute(site),
			routes.MakeImageRoute(site),
			routes.MakeStyleRoute(site),
		},
		Middlewares: []func(http.Handler) http.Handler{
			middleware.RealIP,
			middleware.Logger,
		},
	}
	golam.Run(chi.NewMux(), cfg)
}

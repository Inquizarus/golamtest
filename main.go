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
	site, err := loadSiteConfig()
	nilOrPanic(err)
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

func loadSiteConfig() (config.Site, error) {
	site := config.Site{
		Name: "site",
		Options: config.Options{
			HTTPPort:    "80",
			StaticDir:   "statics",
			TemplateDir: "templates",
		},
	}
	err := config.Load(&site)
	return site, err
}

func nilOrPanic(i interface{}) {
	if nil != i {
		panic(i)
	}
}

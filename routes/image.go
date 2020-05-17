package routes

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/inquizarus/golam"
	"inquizarus.dev/site/config"
)

// MakeImageRoute ...
func MakeImageRoute(site config.Site) golam.Route {
	return golam.Route{
		Pattern: "/img/{path}",
		Get: func(res http.ResponseWriter, req *http.Request) {
			defer req.Body.Close()
			content, err := ioutil.ReadFile(filepath.Join(site.Options.StaticDir, "img", chi.URLParam(req, "path")))
			if nil != err {
				res.WriteHeader(http.StatusNotFound)
				return
			}
			res.Header().Add("content-type", http.DetectContentType(content))
			res.Write(content)
		},
	}
}

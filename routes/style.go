package routes

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/inquizarus/golam"
	"inquizarus.dev/site/config"
)

// MakeStyleRoute ...
func MakeStyleRoute(site config.Site) golam.Route {
	return golam.Route{
		Pattern: "/style/{path}",
		Get: func(res http.ResponseWriter, req *http.Request) {
			defer req.Body.Close()
			content, err := ioutil.ReadFile(filepath.Join(site.Options.StaticDir, "css", chi.URLParam(req, "path")))
			if nil != err {
				res.WriteHeader(http.StatusNotFound)
				return
			}
			res.Header().Add("content-type", "text/css; charset=utf-8")
			res.Write(content)
		},
	}
}

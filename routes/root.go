package routes

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/inquizarus/golam"
	"inquizarus.dev/site/config"
)

// Data ...
type Data struct {
	PageTitle     string
	BodyText      string
	BottomScripts []string
	Styles        []string
}

// MakeRootRoute builds and returns the base routing
// for the test site
func MakeRootRoute(site config.Site) golam.Route {
	return golam.Route{
		Pattern: "/",
		Get: func(res http.ResponseWriter, req *http.Request) {
			defer req.Body.Close()
			tfb, _ := ioutil.ReadFile(filepath.Join(site.Options.TemplateDir, "page"))
			t, _ := template.New("root").Parse(string(tfb))
			t.Execute(res, Data{
				PageTitle: "/",
				BodyText:  "Bacon ipsum dolor amet ground round swine pancetta tongue. Kielbasa andouille venison drumstick t-bone brisket ham ham hock salami.",
				BottomScripts: []string{
					"/js/main.js",
				},
				Styles: []string{
					"https://unpkg.com/modern-css-reset/dist/reset.min.css",
					"/style/main.css",
				},
			})
		},
	}
}

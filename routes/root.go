package routes

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/inquizarus/golam"
	"inquizarus.dev/site/config"
	"inquizarus.dev/site/templates"
)

// Data ...
type Data struct {
	PageTitle     string
	Content       string
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

			// Load the layout
			layoutBytes, _ := ioutil.ReadFile(filepath.Join(site.Options.TemplateDir, templates.Layout.FilePath()))
			layoutTemplate, _ := template.New(templates.Layout.Name()).Parse(string(layoutBytes))

			// Load the page content
			contentBytes, _ := ioutil.ReadFile(filepath.Join(site.Options.TemplateDir, "content", "root.html"))
			contentTemplate, _ := template.New("content").Parse(string(contentBytes))
			contentBuffer := bytes.NewBuffer([]byte{})
			contentTemplate.Execute(contentBuffer, struct{ BodyText string }{""})

			// Inject content into layout
			layoutTemplate.New("content").Parse(string(contentBuffer.Bytes()))

			// Render layout into the response
			layoutTemplate.Execute(res, templates.Layout.Data())
		},
	}
}

package templates

// Layout template for the site
var Layout = StdTemplate{
	templateFile: "layout.html",
	templateName: "layout",
	templateData: struct {
		PageTitle     string
		Content       string
		BottomScripts []string
		Styles        []string
	}{
		PageTitle: "asd",
		BottomScripts: []string{
			"/js/main.js",
		},
		Styles: []string{
			"https://unpkg.com/modern-css-reset/dist/reset.min.css",
			"/style/main.css",
		},
	},
}

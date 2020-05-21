package templates

// Template ...
type Template interface {
	Name() string
	FilePath() string
	Data() interface{}
}

// StdTemplate is not a virus ones
type StdTemplate struct {
	templateName string
	templateFile string
	templateData interface{}
}

// Name returns the string which the template will be registered under
func (t *StdTemplate) Name() string {
	return t.templateName
}

// FilePath generates and return abs path to template file
func (t *StdTemplate) FilePath() string {
	return t.templateFile
}

// Data returns whatever is used to bind content into the template
func (t *StdTemplate) Data() interface{} {
	return t.templateData
}

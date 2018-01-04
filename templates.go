package daytimer

import (
	"fmt"
	"text/template"
)

// LoadTemplate from precompiled binary data with a path relative to the
// templates directory in the project root. Panics on error.
func LoadTemplate(path string) (*template.Template, error) {
	bytes, err := Asset(path)
	if err != nil {
		return nil, fmt.Errorf("unable to parse %s: %s", path, err)
	}

	return template.New(path).Parse(string(bytes))
}

// MustLoadTemplate uses template.Must and panics if an error occurs.
func MustLoadTemplate(path string) *template.Template {
	return template.Must(LoadTemplate(path))
}

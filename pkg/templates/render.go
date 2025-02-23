package templates

import (
	"fmt"
	"net/http"
)

// Renders the page based on the template.
func Render(cache TemplatesCache, w http.ResponseWriter, name string, td *TemplateData) error {
	ts, ok := cache[name]
	if !ok {
		return fmt.Errorf("template %s does not exist", name)
	}
	err := ts.Execute(w, td)
	if err != nil {
		return err
	}
	return nil
}

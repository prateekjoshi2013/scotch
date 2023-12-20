package render

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func (s *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(s.Renderer) {
	case "go":
		return s.GoPage(w, r, view, variables, data)
	case "jet":
	}
	return nil
}

func (s *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", s.RootPath, view))
	if err != nil {
		return err
	}
	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}
	return nil
}

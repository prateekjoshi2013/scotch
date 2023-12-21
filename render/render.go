package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
)

func (s *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(s.Renderer) {
	case "go":
		return s.GoPage(w, r, view, variables, data)
	case "jet":
		return s.JetPage(w, r, view, variables, data)
	default:
		return fmt.Errorf("unsupported renderer: %s", s.Renderer)
	}
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

func (s *Render) JetPage(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	var vars jet.VarMap
	if variables != nil {
		vars = variables.(jet.VarMap)
		} else {
		vars = make(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	t, err := s.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))
	if err != nil {
		return err
	}
	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

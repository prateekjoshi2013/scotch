package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{
		name:          "go_page",
		renderer:      "go",
		template:      "home",
		errorExpected: false,
		errorMessage:  "error rendering go template",
	},
	{
		name:          "go_page_no_template",
		renderer:      "go",
		template:      "no-file",
		errorExpected: true,
		errorMessage:  "no error rendering go template",
	},
	{
		name:          "jet_page",
		renderer:      "jet",
		template:      "home",
		errorExpected: false,
		errorMessage:  "error rendering jet template",
	},
	{
		name:          "jet_page_no_template",
		renderer:      "jet",
		template:      "no-file",
		errorExpected: true,
		errorMessage:  "no error rendering jet template",
	},
	{
		name:          "invalid_renderer",
		renderer:      "foo",
		template:      "home",
		errorExpected: true,
		errorMessage:  "no error rendering with non existent template engine",
	},
}

func TestRender_JetPageNoTemplate(t *testing.T) {

	for _, e := range pageData {

		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}

		// http testing library mocked response writer
		w := httptest.NewRecorder()
		// test renderer rootpath to mocked template pages
		testRenderer.RootPath = "./testdata"

		// test jet template rendering functionality
		testRenderer.Renderer = e.renderer

		// test jet template rendering functionality with no file found
		err = testRenderer.Page(w, r, e.template, nil, nil)

		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}

}

func TestRender__GoPage(t *testing.T) {
	// http testing library mocked response writer
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	// test renderer rootpath to mocked template pages
	testRenderer.RootPath = "./testdata"

	// test go template rendering functionality

	testRenderer.Renderer = "go"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}
}

func TestRender__JetPage(t *testing.T) {
	// http testing library mocked response writer
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	// test go template rendering functionality

	testRenderer.Renderer = "jet"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("error rendering page", err)
	}
}

func TestRender_NoRendererSet(t *testing.T) {
	// http testing library mocked response writer
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	// test renderer rootpath to mocked template pages
	testRenderer.RootPath = "./testdata"

	// test jet template rendering functionality
	testRenderer.Renderer = ""
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err == nil { // should have an error since nofile is not a template
		t.Error("error rendering page", err)
	}
}

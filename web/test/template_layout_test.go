package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/layout.gohtml", "./templates/layout_header.gohtml", "./templates/layout_footer.gohtml"))
	// kenapa tidak layout.gohtml karna sudah diganti namanya dengan {{ define "layout" }}
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template layout",
		"Name":  "Dicki",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

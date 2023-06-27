package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayMyPage(name string) string {
	return "Hello " + name + ", my name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayMyPage "Roy" }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Dicki",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateGlobalFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len .Name }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Dicki",
	})
}

func TestTemplateGlobalFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobalFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateMenambahGlobalFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	// menambahkan global func harus dibuat sebelum di parse
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Dicki",
	})
}

func TestTemplateMenambahGlobalFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateMenambahGlobalFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateGlobalFunctionPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	// menambahkan global func harus dibuat sebelum di parse
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Dicki",
	})
}

func TestTemplateGlobalFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobalFunctionPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

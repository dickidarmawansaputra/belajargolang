package test

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

// menangani jika terjadi panic agar program tidak berhenti & errornya bisa ditampilkan dengan recover
func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("terjadi error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "error: %s", err)
		}
	}()

	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("handler executed")
		fmt.Fprint(writer, "hello middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("foo handler executed")
		fmt.Fprint(writer, "hello foo")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("Oops")
	})

	logMiddleware := LogMiddleware{
		Handler: mux,
	}

	errorHandler := ErrorHandler{
		Handler: &logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

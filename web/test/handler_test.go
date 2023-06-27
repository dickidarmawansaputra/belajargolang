package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	// writer respon yg diberikan ke client
	// request yg dikirim dari client
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello world!")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

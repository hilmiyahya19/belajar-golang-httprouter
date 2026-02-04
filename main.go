package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Hello HTTP Router")
	})
	
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	
	server.ListenAndServe()
}
package main

import (
	"fmt"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/transparentt/go-crud-server/openapi"
)

func main() {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	server := openapi.NewServer()

	router := chi.NewRouter()
	router.Use(middleware.OapiRequestValidator(swagger)) // validation

	openapi.HandlerFromMux(server, router)

	http.ListenAndServe("localhost:8080", router)
}

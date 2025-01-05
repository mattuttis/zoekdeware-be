package main

import (
	_ "embed"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mattuttis/zoekdeware-be/api"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"log"
	"net/http"
	"os"
)

var swaggerSpec []byte

func main() {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	r := mux.NewRouter()

	specRouter := r.PathPrefix("/spec").Subrouter()
	specRouter.PathPrefix("/").Handler(http.StripPrefix("/spec", http.FileServer(http.Dir("api/spec"))))

	swaggerRouter := r.PathPrefix("/swagger-ui").Subrouter()
	swaggerRouter.PathPrefix("/").Handler(http.StripPrefix("/swagger-ui", http.FileServer(http.Dir("static/swagger-ui"))))

	membersRouter := r.PathPrefix("/members").Subrouter()
	membersServer := api.NewServer()
	membersRouter.Use(middleware.OapiRequestValidator(swagger))
	_ = api.HandlerFromMux(membersServer, r)

	statusRouter := r.PathPrefix("/status").Subrouter()
	statusRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	s := &http.Server{
		Handler: handlers.CORS(headersOk, originsOk, methodsOk)(r),
		Addr:    ":8080",
	}

	log.Fatal(s.ListenAndServe())
}

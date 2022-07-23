package main

import (
	"fmt"
	operationService "github.com/submaline/operation-service"
	"github.com/submaline/operation-service/gen/operation/v1/operationv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	operationServiceHandler := &operationService.OperationService{}

	mux := http.NewServeMux()
	mux.Handle(operationv1connect.NewOperationServiceHandler(
		operationServiceHandler))

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Service listening on %v", port)
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

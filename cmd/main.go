package main

import (
	"log"
	"net/http"

	soap "github.com/AnhellO/sample-golang-soap-xml"
)

func main() {
	mux := soap.NewSOAPMux()
	// ts := httptest.NewServer(mux)
	log.Fatal(http.ListenAndServe(":12345", mux))
}

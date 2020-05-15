package main

import (
	"log"
	"net/http"

	"github.com/sebasblancogonz/rest-api-go/pkg/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}

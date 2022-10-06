package main

import (
	"log"

	"github.com/myrachanto/Loader/cmd/routes"
)

func init() {
	log.SetPrefix("tag microservice ")
}
func main() {
	log.Println("Server started")
	routes.ApiLoader()
}

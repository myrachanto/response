package main

import (
	"log"
	"os"

	"github.com/myrachanto/respon/cmd"
	// "github.com/myrachanto/respon/src/routes"
)

func init() {
	log.SetPrefix("tag microservice ")
}
func main() {
	defer os.Exit(0)
	cli := cmd.CommandLine{}
	cli.Run()
	// routes.ApiLoader()
}

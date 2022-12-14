package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/myrachanto/respon/src/api/load"
)

type CommandLine struct {
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" getulr -urls URLS - get sorted urls")
}
func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) getulr(urls string) {
	// fmt.Println("----------------------------step 4")
	if len(urls) == 0 {
		log.Panic("URLS is not Valid")
	}
	// fmt.Println("----------------------------")
	loader := load.NewloadService(load.NewloadRepo())
	res := loader.GetURL(urls)
	// res := load.LoadService.GetURL(urls)
	for _, response := range res {
		fmt.Printf(" %s: has the size of %d \n", response.Url, response.Size)
	}
}
func (cli *CommandLine) Run() {
	cli.validateArgs()

	getURLs := flag.NewFlagSet("getulr", flag.ExitOnError)
	urls := getURLs.String("urls", "", "urls")

	switch os.Args[1] {
	case "getulr":
		err := getURLs.Parse(os.Args[2:])
		Handle(err)
	default:
		runtime.Goexit()
	}

	if getURLs.Parsed() {
		if *urls == "" {
			runtime.Goexit()
		}
		cli.getulr(*urls)
	}

}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

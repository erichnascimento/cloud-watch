package main

import (
	"log"

	"github.com/erichnascimento/cloud-watch/config"
	"github.com/erichnascimento/cloud-watch/watch"
	"github.com/tj/docopt"
	"github.com/tj/go-gracefully"
)

const version = "0.0.1"

const usage = `
	Usage:
	  cloud-watch --config <file>
		cloud-watch --help | -h
		cloud-watch --version | v

	Options:
	  -c, --config  config file to load
	  -h, --help    output help information
	  -v, --version output version

`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)

	if err != nil {
		log.Fatalf("error parsing arguments: %s", err)
	}

	file := args["<file>"].(string)
	c, err := config.New(file)
	if err != nil {
		log.Fatalf("error loading configuration: %s", err)
	}

	log.Printf("starting monitor")
	watch.NewMonitor(c).Start()
	gracefully.Shutdown()
	log.Printf("stopping monitor")
}

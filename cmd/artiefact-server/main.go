package main

import (
	"flag"
	"log"

	artiefact "github.com/hirokazu/artiefact-backend"
)

func main() {
	flag.Parse()
	configFile := flag.Args()

	if len(configFile) == 1 {
		artiefact.Serve(configFile[0])
	} else {
		log.Fatal("configuration file is not specified")
	}
}

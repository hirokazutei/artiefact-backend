package artiefact

import (
	"fmt"
	"log"
)

// Serve serves the server
func Serve(confPath string) {
	fmt.Println(confPath)
	appConfig, err := NewAppConfig(confPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", appConfig)

	database, err := NewDatabase(appConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", database)

}

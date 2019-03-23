package artiefact

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// IapiHandler is a handler
type IapiHandler struct {
	handler func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

// Serve serves the server
func Serve(confPath string) {
	// get the config
	appConfig, err := NewAppConfig(confPath)
	if err != nil {
		log.Fatal(err)
	}

	// start database
	database, err := NewDatabase(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	// set up root context
	app, err := NewApp(appConfig, database)
	if err != nil {
		log.Fatalf("failed to create app: %s", err)
	}

	// application routing
	router := mux.NewRouter()

	// user app
	userApp := &UserApp{app}
	router.HandleFunc("/signup", userApp.SignUpHandler).Methods("POST")

	// listen and serve
	log.Fatal(http.ListenAndServe(":8000", router))
}

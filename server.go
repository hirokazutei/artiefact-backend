package artiefact

import (
	"fmt"
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

	// set up root context
	app, err := NewApp(appConfig, database)
	if err != nil {
		log.Fatalf("failed to create app: %s", err)
	}

	// application routing
	router := mux.NewRouter()

	// user
	userApp := &UserApp{app}
	router.HandleFunc("/signup", userApp.SignUpHandler).Methods("POST")

	// Testing
	router.HandleFunc("/get-user", userApp.GetUserHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

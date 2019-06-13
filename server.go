package artiefact

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/hirokazu/artiefact-backend/constants"
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
		log.Fatalf(c.ErrorActionDetail("creating", "app config", err.Error()))
	}

	// start database
	database, err := NewDatabase(appConfig)
	if err != nil {
		log.Fatalf(c.ErrorActionDetail("starting", "database", err.Error()))
	}

	// set up root context
	app, err := NewApp(appConfig, database)
	if err != nil {
		log.Fatalf(c.ErrorActionDetail("creating", "app", err.Error()))
	}

	// application routing
	router := mux.NewRouter()

	// user app
	userApp := &UserApp{app}
	userRouter := router.PathPrefix("/user").Subrouter()

	// signup
	userRouter.HandleFunc("/signup", userApp.SignUpHandler).Methods("POST")
	userRouter.HandleFunc("/signin", userApp.SignInHandler).Methods("POST")

	// listen and serve
	log.Fatal(http.ListenAndServe(":8000", router))
}

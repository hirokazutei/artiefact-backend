package artiefact

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/hirokazu/artiefact-backend/constants"
	"github.com/justinas/alice"
)

// ServiceHandler handler struct
type ServiceHandler struct {
	handler func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

type contextKeyType int

const (
	contextKeyAuth contextKeyType = iota
)

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

	var serverChain alice.Chain

	middlewareChain := serverChain.Append(
		setJSONHeaderMiddleware,
		tokenAuthMiddleware(app),
		recoverMiddleware,
	)

	// application routing
	router := mux.NewRouter()

	// user app
	userApp := &UserApp{app}
	userRouter := router.PathPrefix("/user").Subrouter()

	// sign-up
	userRouter.Methods("GET").Path("/get-user").Handler(
		middlewareChain.Then(UserHandler{handler: userApp.GetUserHandler}))
	userRouter.Methods("POST").Path("/sign-in").Handler(
		middlewareChain.Then(UserHandler{handler: userApp.SignInHandler}))
	userRouter.Methods("POST").Path("/sign-up").Handler(
		middlewareChain.Then(UserHandler{handler: userApp.SignUpHandler}))
	userRouter.Methods("GET").Path("/username-availability").Handler(
		middlewareChain.Then(UserHandler{handler: userApp.UsernameAvailabilityHandler}))

	// listen and serve
	log.Fatal(http.ListenAndServe(":8000", router))
}

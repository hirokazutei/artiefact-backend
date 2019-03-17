package artiefact

// App base application for all handlers
type App struct {
	Config *AppConfig
	DB     DBer
}

// NewApp create new app
func NewApp(c *AppConfig, db DBer) (*App, error) {
	app := &App{
		Config: c,
		DB:     db,
	}
	return app, nil
}

package main

func (app *Application) Routes() {
	app.server.GET("/", app.handler.HealthCheck)
}

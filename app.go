package nifty

import "fmt"

type App struct {
	ctrl controller
	Cfg  Config
}

func CreateApp(cfg Config, urls UrlMapper) App {
	var app App
	app.ctrl = createController(urls)
	app.Cfg = cfg
	return app
}

func (a App) Start() {
	fmt.Println("app started listening")
	a.ctrl.Listen(a.Cfg.Port)
}

package nifty

type App struct {
	ctrl controller
	Cfg  Config
}

func CreateApp(cfg Config, urls UrlMapper) App {
	var app App
	app.ctrl = createController(urls)

	return app
}

func (a App) Start() {
	a.ctrl.Listen(a.Cfg.Port)
}

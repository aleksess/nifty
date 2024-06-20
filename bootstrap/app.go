package bootstrap

import (
	"github.com/aleksess/nifty"
	"github.com/aleksess/nifty/bootstrap/config"
)

var App = nifty.CreateApp(config.Config, Urls)

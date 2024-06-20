package bootstrap

import (
	"github.com/aleksess/nifty"
	"github.com/aleksess/nifty/bootstrap/views"
)

var Urls = nifty.UrlMapper{
	nifty.UrlMapping{View: views.Hello, Method: nifty.Get, Url: "/"},
}

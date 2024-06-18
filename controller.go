package nifty

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	router chi.Router
}

func NewController() Controller {
	return Controller{router: chi.NewRouter()}
}

func (c Controller) SetUrls(u []UrlMapping) {

	for _, m := range u {
		switch m.Method {
		case Post:
			c.router.With(m.Middleware...).Post(m.Url, http.HandlerFunc(m.View))
		case Put:
			c.router.With(m.Middleware...).Put(m.Url, http.HandlerFunc(m.View))
		case Patch:
			c.router.With(m.Middleware...).Patch(m.Url, http.HandlerFunc(m.View))
		case Delete:
			c.router.With(m.Middleware...).Delete(m.Url, http.HandlerFunc(m.View))
		case Options:
			c.router.With(m.Middleware...).Options(m.Url, http.HandlerFunc(m.View))
		case Head:
			c.router.With(m.Middleware...).Head(m.Url, http.HandlerFunc(m.View))
		case Connect:
			c.router.With(m.Middleware...).Connect(m.Url, http.HandlerFunc(m.View))
		case Trace:
			c.router.With(m.Middleware...).Trace(m.Url, http.HandlerFunc(m.View))
		case Get:
			fallthrough
		default:
			c.router.With(m.Middleware...).Get(m.Url, http.HandlerFunc(m.View))

		}

	}
}

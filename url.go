package nifty

import "net/http"

const (
	Get = iota
	Post
	Put
	Patch
	Delete
	Options
	Head
	Connect
	Trace
)

type UrlMapping struct {
	Method     int
	Url        string
	View       View
	Middleware []func(http.Handler) http.Handler
}

type UrlMapper []UrlMapping

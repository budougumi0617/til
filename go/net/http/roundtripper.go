package http

import (
	stdhttp "net/http"
)

type RoundTripper interface {
	RoundTrip(*stdhttp.Request) (*stdhttp.Response, error)
}

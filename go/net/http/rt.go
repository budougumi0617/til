package http

import (
	ghttp "net/http"
)

type transport struct {
	Base ghttp.RoundTripper
}

func (t *transport) base() ghttp.RoundTripper {
	if t.Base == nil {
		return ghttp.DefaultTransport
	}
	return t.Base
}

// RoundTrip is custom RoundTrip.
func (t *transport) RoundTrip(req *ghttp.Request) (*ghttp.Response, error) {
	// do anything...
	return t.base().RoundTrip(req)
}

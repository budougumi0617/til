package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

var myt = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 3 * time.Second,
		DualStack: true,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	// DisableKeepAlives:     true,
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cli := &http.Client{
		Transport: myt,
	}
	var result httpstat.Result
	req, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodGet, "https://budougumi0617.github.io",
		nil,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := cli.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	result.End(time.Now())

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "response code: %+v", result)
}

func main() {
	serverAddress := ":8080"
	srv := &http.Server{
		Addr:    serverAddress,
		Handler: http.HandlerFunc(indexHandler),
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

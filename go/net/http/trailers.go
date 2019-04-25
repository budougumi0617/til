package http

import (
	"io"
	"net/http"
)

func start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/sendstrailers", func(w http.ResponseWriter, req *http.Request) {
		// Before any call to WriteHeader or Write, declare
		// the trailers you will set during the HTTP
		// response. These three headers are actually sent in
		// the trailer.
		w.Header().Set("Trailer", "AtEnd1, AtEnd2")
		w.Header().Add("Trailer", "AtEnd3")

		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)

		w.Header().Set("AtEnd1", "value 1") // Trailer because set after WriteHeader.
		io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
		w.Header().Set("AtEnd2", "value 2")
		w.Header().Set("AtEnd3", "value 3") // These will appear as trailers.
	})
	http.ListenAndServe(":8080", mux)
}

// curl -v -X GET http://localhost:8080/sendstrailers
// Note: Unnecessary use of -X or --request, GET is already inferred.
// *   Trying ::1...
// * TCP_NODELAY set
// * Connected to localhost (::1) port 8080 (#0)
// > GET /sendstrailers HTTP/1.1
// > Host: localhost:8080
// > User-Agent: curl/7.54.0
// > Accept: */*
// >
// < HTTP/1.1 200 OK
// < Content-Type: text/plain; charset=utf-8
// < Trailer: AtEnd1, AtEnd2
// < Trailer: AtEnd3
// < Date: Sun, 03 Mar 2019 22:25:51 GMT
// < Transfer-Encoding: chunked
// <
// This HTTP response has both headers before this text and trailers at the end.
// * Connection #0 to host localhost left intact

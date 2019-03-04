# How to send cookies by curl
Use `-b` option. If know cookies in response, you are able to check use by `-i` option.

```bash
$ curl -i -b "hoge=bar; test=value" -X GET localhost:8080

HTTP/1.1 200 OK
Date: Sun, 03 Mar 2019 23:32:21 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Hello, World%
```

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for i, c := range r.Cookies() {
		fmt.Printf("cookie[%d] %s:%s\n", i, c.Name, c.Value)
		// http.SetCookie(w, c)
	}
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.ListenAndServe(":8080", r)
}
```

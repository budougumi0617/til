FROM golang:1.15.6-alpine3.12 as build-env

ENV CGO_ENABLED 0

RUN apk add --no-cache git

WORKDIR /debuggingTutorial/
ADD . /debuggingTutorial/

RUN go build -o /debuggingTutorial/srv .

# Get Delve from a GOPATH not from a Go Modules project
WORKDIR /go/src/
RUN go get github.com/go-delve/delve/cmd/dlv

FROM alpine:3.12 as debugger

WORKDIR /
COPY --from=build-env /debuggingTutorial/srv /
COPY --from=build-env /go/bin/dlv /

EXPOSE 8080 40000

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/srv"]

FROM alpine:3.12 as server
COPY --from=build-env /debuggingTutorial/srv /
EXPOSE 8080
CMD ["/srv"]

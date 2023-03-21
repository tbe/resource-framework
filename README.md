# DEPRECTAION NOTICE

This repostiry is deprecated. Please move to https://github.com/tbe/go-concourse-resource

# resource-framework

*a GoLang framework to implement [concourse-ci][1] resources*

## Usage

For full usage, please see [godoc][2]

```go
package main

import (
	"log"

	"github.com/tbe/resource-framework/resource"
)

// DummySource defines the source configuration for our dummy resource
type DummySource struct {
	DoSomething bool `json:"do_something"`
}

type DummyVersion struct {
	Counter int `json:"counter" validate:"gt=0"`
}

// DummyResource implements a dummy check resource
type DummyResource struct {
	source  *DummySource
	version *DummyVersion
}

func NewDummyResource() *DummyResource {
	return &DummyResource{
		source:  &DummySource{},
		version: &DummyVersion{},
	}
}

func (r *DummyResource) Source() (source interface{}) {
	return r.source
}

func (r *DummyResource) Version() (version interface{}) {
	return r.version
}

func (r *DummyResource) Check() (version interface{}, err error) {
	if r.source.DoSomething {
		r.version.Counter++
	}
	return r.version, nil
}

func main() {
	r := NewDummyResource()

	handler, err := resource.NewHandler(r)
	if err != nil {
		log.Fatal(err)
	}
	_ = handler.Run()
}
```

## Testing

The `test` package provides a helper for testing resources, written with the `resource-framwork`.
Have a look at the tests for the `resource` package and on [godoc][3]

## Resource Building

To build a docker container for your resource, you could to this with a Dockerfile like this:

```Dockerfile
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o resource .
RUN mkdir -p /target/opt/resource/
RUN cp resource /target/opt/resource/
RUN ln -s resource /target/opt/resource/in
RUN ln -s resource /target/opt/resource/out
RUN ln -s resource /target/opt/resource/check


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /target/opt /opt
```

[1]: https://concourse-ci.org/
[2]: https://godoc.org/github.com/tbe/resource-framework/resource
[3]: https://godoc.org/github.com/tbe/resource-framework/test

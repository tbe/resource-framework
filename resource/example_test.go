package resource_test

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

func ExampleMain() {
	r := NewDummyResource()

	handler, err := resource.NewHandler(r)
	if err != nil {
		log.Fatal(err)
	}
	_ = handler.Run()
}

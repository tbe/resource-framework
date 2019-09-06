package resource_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/test"
)

type TestCheckSource struct {
	Some    string `json:"some"`
	Testing bool   `json:"testing"`
}

type TestCheckVersion struct {
	Number int `json:"number"`
}

type testCheckResource struct {
	source  *TestCheckSource
	version *TestCheckVersion
}

func (o *testCheckResource) Source() interface{} {
	o.source = &TestCheckSource{}
	return o.source
}

func (o *testCheckResource) Version() interface{} {
	o.version = &TestCheckVersion{}
	return o.version
}

func (o *testCheckResource) Check() (interface{}, error) {
	return o.version, nil
}

func TestHandler_Check(t *testing.T) {
	res := &testCheckResource{}
	testHandler := test.NewHandler(t,res)

	// verify the test call
	testHandler.TestCheck(
		`{"source":{"some":"string","testing":true},"version":{"number": 42}}`,
		`{"number":42}`,
	)

	// verify the input parsing
	assert.Equal(t, "string", res.source.Some)
	assert.Equal(t, true, res.source.Testing)
	assert.Equal(t, 42, res.version.Number)
}

package resource_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/test"
)

type TestOutSource struct {
	Some    string `json:"some"`
	Testing bool   `json:"testing"`
}

type TestOutVersion struct {
	Number int `json:"number"`
}

type TestOutParams struct {
	Why    string `json:"why"`
	Really bool   `json:"really"`
}

type TestOutMetadata struct {
	Status string `json:"status"`
}

type testOutResource struct {
	source *TestOutSource
	params *TestOutParams
}

func (o *testOutResource) Source() (source interface{}) {
	o.source = &TestOutSource{}
	return o.source
}

func (o *testOutResource) Params() (version interface{}) {
	o.params = &TestOutParams{}
	return o.params
}

func (o *testOutResource) Out(dir string) (version interface{}, metadata []interface{}, err error) {
	return &TestOutVersion{Number: 42}, []interface{}{&TestOutMetadata{Status: "cool"}, &TestOutMetadata{Status: "really cool"}}, nil
}

func TestHandler_Out(t *testing.T) {
	res := &testOutResource{}

	handler := test.NewHandler(t, res)

	handler.TestOut(
		`{"source":{"some":"string","testing":true},"params":{"why": "because", "really": true}}`,
		`{"metadata":[{"status":"cool"},{"status":"really cool"}],"version":{"number":42}}`,
	)

	// verify the parser result
	assert.Equal(t, "string", res.source.Some)
	assert.Equal(t, true, res.source.Testing)
	assert.Equal(t, "because", res.params.Why)
	assert.Equal(t, true, res.params.Really)
}

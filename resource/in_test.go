package resource_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/test"
)

type TestInSource struct {
	Some    string `json:"some"`
	Testing bool   `json:"testing"`
}

type TestInVersion struct {
	Number int `json:"number"`
}

type TestInMetadata struct {
	Status string `json:"status"`
}

type testInResource struct {
	source  *TestInSource
	version *TestInVersion
}

func (i *testInResource) Source() (source interface{}) {
	i.source = &TestInSource{}
	return i.source
}

func (i *testInResource) Version() (version interface{}) {
	i.version = &TestInVersion{}
	return i.version
}

func (i *testInResource) In(dir string) (version interface{}, metadata []interface{}, err error) {
	return i.version, []interface{}{&TestInMetadata{Status: "cool"}, &TestInMetadata{Status: "really cool"}}, nil
}

func TestHandler_In(t *testing.T) {
	res := &testInResource{}
	handler := test.NewHandler(t, res)

	handler.TestIn(
		`{"source":{"some":"string","testing":true},"version":{"number": 42}}`,
		`{"metadata":[{"status":"cool"},{"status":"really cool"}],"version":{"number":42}}`,
	)

	// verify the parser result
	assert.Equal(t, "string", res.source.Some)
	assert.Equal(t, true, res.source.Testing)
	assert.Equal(t, 42, res.version.Number)
}

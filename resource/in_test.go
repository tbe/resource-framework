package resource_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/resource"
	"github.com/tbe/resource-framework/test"
)

type TestInSource struct {
	Some    string `json:"some" validate:"required"`
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
	test.AutoTestIn(t, func() resource.Resource {
		return &testInResource{}
	}, test.CaseList{
		"valid input": {
			Input:  `{"source":{"some":"string","testing":true},"version":{"number": 42}}`,
			Output: `{"metadata":[{"status":"cool"},{"status":"really cool"}],"version":{"number":42}}`,
			Validation: func(assertion *assert.Assertions, res interface{}) {
				r := res.(*testInResource)
				assertion.Equal("string", r.source.Some)
				assertion.Equal(true, r.source.Testing)
				assertion.Equal(42, r.version.Number)
			},
		},
		"invalid input": {
			Input:       `{"source":{"testing":true},"version":{"number": 42}}`,
			ShouldFail:  true,
			ErrorString: "failed to validate input: Key: 'inInput.Source.Some' Error:Field validation for 'Some' failed on the 'required' tag",
		},
	})
}

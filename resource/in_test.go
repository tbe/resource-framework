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

type TestInParams struct {
	Another string `json:"another" validate:"required"`
	Param   int    `json:"param"`
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
	params  *TestInParams
}

func (i *testInResource) Source() (source interface{}) {
	i.source = &TestInSource{}
	return i.source
}

func (i *testInResource) Version() (version interface{}) {
	i.version = &TestInVersion{}
	return i.version
}

func (i *testInResource) Params() (version interface{}) {
	i.params = &TestInParams{}
	return i.params
}

func (i *testInResource) In(_ string) (version interface{}, metadata []interface{}, err error) {
	return i.version, []interface{}{&TestInMetadata{Status: "cool"}, &TestInMetadata{Status: "really cool"}}, nil
}

func TestHandler_In(t *testing.T) {
	test.AutoTestIn(t, func() resource.Resource {
		return &testInResource{}
	}, test.CaseList{
		"valid input": {
			Input:  `{"source":{"some":"string","testing":true},"version":{"number": 42},"params":{"another":"thing","param":23}}`,
			Output: `{"metadata":[{"status":"cool"},{"status":"really cool"}],"version":{"number":42}}`,
			Validation: func(_ *testing.T, assertion *assert.Assertions, res interface{}) {
				r := res.(*testInResource)
				assertion.Equal("string", r.source.Some)
				assertion.Equal(true, r.source.Testing)
				assertion.Equal(42, r.version.Number)
				assertion.Equal("thing", r.params.Another)
				assertion.Equal(23, r.params.Param)
			},
		},
		"invalid input": {
			Input:       `{"source":{"testing":true},"version":{"number": 42},"params":{"another":"thing","param":23}}`,
			ShouldFail:  true,
			ErrorString: "failed to validate input: Key: 'inInput.Source.Some' Error:Field validation for 'Some' failed on the 'required' tag",
		},
	})
}

package resource_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/resource"
	"github.com/tbe/resource-framework/test"
)

type TestCheckSource struct {
	Some    string `json:"some" validate:"required"`
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
	test.AutoTestCheck(t, func() resource.Resource { return &testCheckResource{} }, map[string]test.Case{
		"valid input": {
			Input:  `{"source":{"some":"string","testing":true},"version":{"number": 42}}`,
			Output: `{"number":42}`,
			Validation: func(assertion *assert.Assertions, res interface{}) {
				r := res.(*testCheckResource)
				assertion.Equal("string", r.source.Some)
				assertion.Equal(true, r.source.Testing)
				assertion.Equal( 42, r.version.Number)
			},
		},
		"invalid input": {
			Input:       `{"source":{"testing":true},"version":{"number": 42}}`,
			ShouldFail:  true,
			ErrorString: "failed to validate input: Key: 'checkInput.Source.Some' Error:Field validation for 'Some' failed on the 'required' tag",
		},
	})
}

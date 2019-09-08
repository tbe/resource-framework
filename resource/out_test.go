package resource_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/resource"
	"github.com/tbe/resource-framework/test"
)

type TestOutSource struct {
	Some    string `json:"some" validate:"required"`
	Testing bool   `json:"testing"`
}

type TestOutVersion struct {
	Number int `json:"number"`
}

type TestOutParams struct {
	Why    string `json:"why"`
	Really bool   `json:"really"`
	Count  int    `json:"count" validate:"gt=1"`
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
	test.AutoTestOut(t, func() resource.Resource {
		return &testOutResource{}
	}, test.CaseList{
		"valid input": {
			Input:  `{"source":{"some":"string","testing":true},"params":{"why": "because", "really": true, "count": 2}}`,
			Output: `{"metadata":[{"status":"cool"},{"status":"really cool"}],"version":{"number":42}}`,
			Validation: func(assertions *assert.Assertions, res interface{}) {
				r := res.(*testOutResource)
				assertions.Equal("string", r.source.Some)
				assertions.Equal(true, r.source.Testing)
				assertions.Equal("because", r.params.Why)
				assertions.Equal(true, r.params.Really)
			},
		},
		"incomplete source": {
			Input:       `{"source":{"testing":true},"params":{"why": "because", "really": true,"count": 2}}`,
			ShouldFail:  true,
			ErrorString: "failed to validate input: Key: 'outInput.Source.Some' Error:Field validation for 'Some' failed on the 'required' tag",
		},
		"invalid params": {
			Input:       `{"source":{"some":"string","testing":true},"params":{"why": "because", "really": false,"count":0}}`,
			ShouldFail:  true,
			ErrorString: "failed to validate input: Key: 'outInput.Params.Count' Error:Field validation for 'Count' failed on the 'gt' tag",
		},
	})
	// verify the parser result

}

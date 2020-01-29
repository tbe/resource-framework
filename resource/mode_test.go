package resource_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/resource"
	"github.com/tbe/resource-framework/test"
)

type TestModeSource struct {
	Some    string `json:"some" validate:"required"`
	Testing bool   `json:"testing"`
}

type TestModeInParams struct {
	Another string `json:"another" validate:"required"`
	Param   int    `json:"param"`
}

type TestModeOutParams struct {
	TestModeInParams
	Extra string `json:"extra" validate:"required"`
}

type TestModeVersion struct {
	Number int `json:"number"`
}

type TestModeMetadata struct {
	Status string `json:"status"`
}

type testModeResource struct {
	source    *TestModeSource
	version   *TestModeVersion
	inParams  *TestModeInParams
	outParams *TestModeOutParams

	mode resource.ResourceMode
}

func (t *testModeResource) SetMode(mode resource.ResourceMode) {
	t.mode = mode
}

func (t *testModeResource) Source() (source interface{}) {
	t.source = &TestModeSource{}
	return t.source
}

func (t *testModeResource) Version() (version interface{}) {
	t.version = &TestModeVersion{}
	return t.version
}

func (t *testModeResource) Params() (params interface{}) {
	switch t.mode {
	case resource.IN:
		t.inParams = &TestModeInParams{}
		return t.inParams
	case resource.OUT:
		t.outParams = &TestModeOutParams{}
		return t.outParams
	}
	return nil
}

func (t *testModeResource) In(_ string) (version interface{}, metadata []interface{}, err error) {
	return t.version, []interface{}{&TestModeMetadata{Status: "cool"}, &TestModeMetadata{Status: "really cool"}}, nil
}

func (t *testModeResource) Out(dir string) (version interface{}, metadata []interface{}, err error) {
	return &TestModeVersion{Number: 42}, []interface{}{&TestModeMetadata{Status: "cool"}, &TestModeMetadata{Status: "really cool"}}, nil
}

func TestHandler_Mode(t *testing.T) {
	test.AutoTestIn(t, func() resource.Resource {
		return &testModeResource{}
	}, test.CaseList{
		"valid input": {
			Input:  `{"source":{"some":"string","testing":true},"version":{"number": 42},"params":{"another":"thing","param":23}}`,
			Output: `{"metadata":[{"status":"cool"},{"status":"really cool"}],"version":{"number":42}}`,
			Validation: func(_ *testing.T, assertion *assert.Assertions, res interface{}) {
				r := res.(*testModeResource)
				assertion.Equal("string", r.source.Some)
				assertion.Equal(true, r.source.Testing)
				assertion.Equal(42, r.version.Number)
				assertion.NotNil(r.inParams)
				assertion.Equal("thing", r.inParams.Another)
				assertion.Equal(23, r.inParams.Param)
			},
		},
	})

	test.AutoTestOut(t, func() resource.Resource {
		return &testModeResource{}
	}, test.CaseList{
		"valid output": {
			Input:  `{"source":{"some":"string","testing":true},"params":{"another":"thing","param":23,"extra":"args"}}`,
			Output: `{"metadata":[{"status":"cool"},{"status":"really cool"}],"version":{"number":42}}`,
			Validation: func(_ *testing.T, assertion *assert.Assertions, res interface{}) {
				r := res.(*testModeResource)
				assertion.Equal("string", r.source.Some)
				assertion.Equal(true, r.source.Testing)
				assertion.Equal("thing", r.outParams.Another)
				assertion.Equal(23, r.outParams.Param)
				assertion.Equal("args", r.outParams.Extra)
			},
		},
		"invalid output": {
			Input:       `{"source":{"some":"string","testing":true},"version":{"number": 42},"params":{"another":"thing","param":23}}`,
			ShouldFail:  true,
			ErrorString: "failed to validate input: Key: 'outInput.Params.Extra' Error:Field validation for 'Extra' failed on the 'required' tag",
		},
	})
}

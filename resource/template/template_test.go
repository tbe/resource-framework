package template_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/resource/template"
	// import our test package, so it plays with the Get* wrappers
	_ "github.com/tbe/resource-framework/test"
)

func TestBuildURL(t *testing.T) {
	assert.Equal(t, "http://localhost/teams/main/pipelines/testpipeline/jobs/testjob/builds/42", template.BuildURL())
}

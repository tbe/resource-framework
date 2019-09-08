package template

import (
	"os"
)

// GetExternalURL holds a function get the contents of ATC_EXTERNAL_URL
var GetExternalURL func() string

// GetTeamName holds a function get the contents of BUILD_TEAM_NAME
var GetTeamName func() string

// GetPipelineName holds a function get the contents of BUILD_PIPELINE_NAME
var GetPipelineName func() string

// GetJobName holds a function get the contents of BUILD_JOB_NAME
var GetJobName func() string

// GetBuildName holds a function get the contents of BUILD_NAME
var GetBuildName func() string

// BuildID holds a function get the contents of BUILD_ID
var GetBuildID func() string

// set our default values for the getter functions
func init() {
	GetExternalURL = func() string {
		return os.Getenv("ATC_EXTERNAL_URL")
	}

	GetTeamName = func() string {
		return os.Getenv("BUILD_TEAM_NAME")
	}

	GetPipelineName = func() string {
		return os.Getenv("BUILD_PIPELINE_NAME")
	}

	GetJobName = func() string {
		return os.Getenv("BUILD_JOB_NAME")
	}

	GetBuildName = func() string {
		return os.Getenv("BUILD_NAME")
	}

	GetBuildID = func() string {
		return os.Getenv("BUILD_ID")
	}
}

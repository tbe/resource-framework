package template

import (
	"os"
)

// ShellExpand is a replacement for os.ExpandEnv, handling only the known variables exported
// by concourse. The main reason to do this, is: testing. As a side effect, this can prevent information leaks
func ShellExpand(input string) string {
	return os.Expand(input, func(s string) string {
		switch s {
		case "ATC_EXTERNAL_URL":
			return GetExternalURL()
		case "BUILD_TEAM_NAME":
			return GetTeamName()
		case "BUILD_PIPELINE_NAME":
			return GetPipelineName()
		case "BUILD_JOB_NAME":
			return GetJobName()
		case "BUILD_NAME":
			 return  GetBuildName()
		case "BUILD_ID":
			return GetBuildID()
		}
		return ""
	})
}

package template

// BuildURL returns the full URL to the current build
func BuildURL() string {
	return ShellExpand("${ATC_EXTERNAL_URL}/teams/${BUILD_TEAM_NAME}/pipelines/${BUILD_PIPELINE_NAME}" +
		"/jobs/${BUILD_JOB_NAME}/builds/${BUILD_NAME}")
}

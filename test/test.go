/*
package test provides a testing framework for resources
*/
package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/log"
	"github.com/tbe/resource-framework/resource"
)

var Log *Logger

// as soon as the testing library is loaded, we replace our logger
func init() {
	Log = NewLogger()
	log.Log = Log
}

// ResourceFactory is a function that returns a new, initialized, resource
type ResourceFactory func() resource.Resource

// Case defines the input and expected result of a single test
type Case struct {
	Input       string                                           // the input for the test
	Output      string                                           // the Output for the test. Can be empty
	ShouldFail  bool                                             // defines if the testcase should fail
	ErrorString string                                           // an optional error message to with
	Validation  func(assertions *assert.Assertions, res interface{}) // an optional validation function
}

// CaseList is a list of named test cases
type CaseList map[string]Case

// Handler is a wrapper around `resource.Handler` with test specific functions
type Handler struct {
	inbuf      *bytes.Buffer
	outbuf     *bytes.Buffer
	reshandler *resource.Handler
	assert     *assert.Assertions
}

// NewHandler returns a new testing handler
func NewHandler(t *testing.T, r interface{}) *Handler {
	rh, err := resource.NewHandler(r)
	if err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		inbuf:      new(bytes.Buffer),
		outbuf:     new(bytes.Buffer),
		assert:     assert.New(t),
		reshandler: rh,
	}

	h.reshandler = rh
	rh.Stdin = h.inbuf
	rh.Stdout = h.outbuf

	return h
}

// Reset clears the internal buffers
func (h *Handler) Reset() {
	h.inbuf.Reset()
	h.outbuf.Reset()
}

func (h *Handler) verifyOutput(expected string, msgAndArgs ...interface{}) bool {
	return h.assert.JSONEq(expected, h.outbuf.String(), msgAndArgs)
}

// LogContains verifies that the log output contains the specified string for the given level.
// Only whole log messages are matched.
func (h *Handler) LogContains(level string, message string) bool {
	return h.assert.Contains(Log.messages[level], message)
}

// NoErrors is a helper to verify that we didn't log any errors.
// This function can be used if you have cases where the resource can complete even if there are errors
func (h *Handler) NoErrors() bool {
	return h.assert.Empty(Log.messages["error"])
}

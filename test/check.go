package test

import (
	"os"
	"testing"
)

// AutoTestCheck is a wrapper around the testing Handler to allow automatic testing
// of multiple test cases for Check resources
func AutoTestCheck(t *testing.T, factory ResourceFactory, cases CaseList) {
	for name, c := range cases {
		res := factory()
		t.Run(name, func(t *testing.T) {
			h := NewHandler(t, res)
			success := false
			if !c.ShouldFail {
				success = h.TestCheck(c.Input, c.Output)
			} else {
				if c.ErrorString != "" {
					success = h.TestCheckShouldFailWithError(c.Input, c.ErrorString)
				} else {
					success = h.TestCheckShouldFail(c.Input)
				}
			}
			if success && c.Validation != nil {
				c.Validation(t, h.assert, res)
			}
		})
	}
}

func (h *Handler) runCheck(input string, f func(err error) bool) bool {
	// clear the logger
	Log.Reset()
	defer h.Reset()

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"check"}

	h.inbuf.Write([]byte(input))
	return f(h.reshandler.Run())
}

// TestCheck verifies the Check action in and output
func (h *Handler) TestCheck(input string, output string) bool {
	return h.runCheck(input, func(err error) bool {
		if !h.assert.NoError(err) {
			return false
		}
		return h.verifyOutput(output)

	})
}

// TestCheckShouldFailWithError runs the Check action and verifies that it failed with the given error string
func (h *Handler) TestCheckShouldFailWithError(input string, errorString string) bool {
	return h.runCheck(input, func(err error) bool {
		return h.assert.EqualError(err, errorString)
	})
}

// TestCheckShouldFail runs the Check action and verifies that it failed
func (h *Handler) TestCheckShouldFail(input string) bool {
	return h.runCheck(input, func(err error) bool {
		return h.assert.Error(err)
	})
}

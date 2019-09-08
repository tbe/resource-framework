package test

import (
	"os"
	"testing"
)

// AutoTestOut is a wrapper around the testing Handler to allow automatic testing
// of multiple test cases for Out resources
func AutoTestOut(t *testing.T, factory ResourceFactory, cases CaseList) {
	for name, c := range cases {
		res := factory()
		t.Run(name, func(t *testing.T) {
			h := NewHandler(t, res)
			success := false
			if !c.ShouldFail {
				success = h.TestOut(c.Input, c.Output)
			} else {
				if c.ErrorString != "" {
					success = h.TestOutShouldFailWithError(c.Input, c.ErrorString)
				} else {
					success = h.TestOutShouldFail(c.Input)
				}
			}
			if success && c.Validation != nil {
				c.Validation(h.assert, res)
			}
		})
	}
}

func (h *Handler) runOut(input string, f func(err error) bool) bool {
	// clear the logger
	Log.Reset()
	defer h.Reset()

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"out", "workdir"}

	h.inbuf.Write([]byte(input))
	return f(h.reshandler.Run())
}

// TestOut verifies the Out action and its output
func (h *Handler) TestOut(input string, output string) bool {
	return h.runOut(input, func(err error) bool {
		if !h.assert.NoError(err) {
			return false
		}
		return h.verifyOutput(output)
	})
}

// TestOutShouldFailWithError runs the Out action and verifies that it failed with the given error string
func (h *Handler) TestOutShouldFailWithError(input string, errorString string) bool {
	return h.runOut(input, func(err error) bool {
		return h.assert.EqualError(err, errorString)
	})
}

// TestOutShouldFail runs the Out action and verifies that it failed
func (h *Handler) TestOutShouldFail(input string) bool {
	return h.runOut(input, func(err error) bool {
		return h.assert.Error(err)
	})
}

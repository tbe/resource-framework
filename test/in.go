package test

import (
	"os"
	"testing"
)

// AutoTestIn is a wrapper around the testing Handler to allow automatic testing
// of multiple test cases for In resources
func AutoTestIn(t *testing.T, factory ResourceFactory, cases CaseList) {
	for name, c := range cases {
		res := factory()
		t.Run(name, func(t *testing.T) {
			h := NewHandler(t, res)
			success := false
			if !c.ShouldFail {
				success = h.TestIn(c.Input, c.Output)
			} else {
				if c.ErrorString != "" {
					success = h.TestInShouldFailWithError(c.Input, c.ErrorString)
				} else {
					success = h.TestInShouldFail(c.Input)
				}
			}
			if success && c.Validation != nil {
				c.Validation(t, h.assert, res)
			}
		})
	}
}

func (h *Handler) runIn(input string, f func(err error) bool) bool {
	// clear the logger
	Log.Reset()
	defer h.Reset()

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"in", "workdir"}

	h.inbuf.Write([]byte(input))
	return f(h.reshandler.Run())
}

// TestIn verifies the In action and output
func (h *Handler) TestIn(input string, output string) bool {
	return h.runIn(input, func(err error) bool {
		if !h.assert.NoError(err) {
			return false
		}
		return h.verifyOutput(output)
	})
}

// TestInShouldFailWithError runs the In action and verifies that it failed with the given error string
func (h *Handler) TestInShouldFailWithError(input string, errorString string) bool {
	return h.runIn(input, func(err error) bool {
		return h.assert.EqualError(err, errorString)
	})
}

// TestInShouldFail runs the In action and verifies that it failed
func (h *Handler) TestInShouldFail(input string) bool {
	return h.runIn(input, func(err error) bool {
		return h.assert.Error(err)
	})
}

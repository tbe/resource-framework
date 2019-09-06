/*
package test provides a testing framework for resources
 */
package test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tbe/resource-framework/resource"
)

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

// TestCheck verifies the Check action
func (h *Handler) TestCheck(input string, output string) bool {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"check"}

	h.inbuf.Write([]byte(input))
	h.reshandler.Run()
	result := h.verifyOutput(output)
	h.Reset()
	return result
}

// TestIn verifies the In action
func (h *Handler) TestIn(input string, output string) bool {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"in", "workdir"}

	h.inbuf.Write([]byte(input))
	h.reshandler.Run()
	result := h.verifyOutput(output)
	h.Reset()
	return result
}

// TestOut verifies the out action
func (h *Handler) TestOut(input string, output string) bool {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"out", "workdir"}

	h.inbuf.Write([]byte(input))
	h.reshandler.Run()
	result := h.verifyOutput(output)
	h.Reset()
	return result
}

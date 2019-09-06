package resource

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path"

	"github.com/tbe/resource-framework/log"
)

// The Handler takes care about all communication with concourse.
type Handler struct {
	check CheckResource
	in    InResource
	out   OutResource

	// Stdout is exported to allow mocking of the communication with concourse
	Stdout io.Writer
	// Stdin is exported to allow mocking of the communication with concourse
	Stdin io.Reader
}

// NewHandler creates a new handler for the given resource
func NewHandler(resource interface{}) (*Handler, error) {
	h := &Handler{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
	}
	valid := false
	if check, ok := resource.(CheckResource); ok {
		h.check = check
		valid = true
	}
	if in, ok := resource.(InResource); ok {
		h.in = in
		valid = true
	}
	if out, ok := resource.(OutResource); ok {
		h.out = out
		valid = true
	}
	if !valid {
		return nil, errors.New("not a valid resource")
	}

	return h, nil
}

// Run checks the name of the called binary and executes the corresponding handler
func (h *Handler) Run() {
	switch path.Base(os.Args[0]) {
	case "check":
		h.Check()
	case "in":
		h.In()
	case "out":
		h.Out()
	default:
		log.Error("unknown action %v", os.Args[0])
	}
}

func (h *Handler) output(data interface{}) error {
	return json.NewEncoder(h.Stdout).Encode(data)
}

func (h *Handler) input(data interface{}) error {
	return json.NewDecoder(h.Stdin).Decode(data)
}

package resource

import (
	"os"

	ji "github.com/tbe/resource-framework/internal/jsoninterface"
	"github.com/tbe/resource-framework/log"
)

type inInput struct {
	Source  *ji.Interface `json:"source"`
	Version *ji.Interface `json:"version"`
}

type inOutput struct {
	Metadata []interface{} `json:"metadata"`
	Version  interface{}   `json:"version"`
}

// In calls the In function of the resource implementation and handles the communication with concourse
func (h *Handler) In() {
	if len(os.Args) < 2 {
		log.Error("missing commandline argument")
		return
	}

	// check for a valid in resource
	if h.in == nil {
		_ = h.output(struct{}{})
		return
	}

	// get the storage for the source
	source := h.in.Source()
	if err := validateStructPtr(source); err != nil {
		log.Error("invalid source storage: %v", err)
	}

	// get the storage for the version
	version := h.in.Version()
	if err := validateStructPtr(version); err != nil {
		log.Error("invalid version storage: %v", err)
	}

	input := &inInput{
		Source:  ji.NewInterface(source),
		Version: ji.NewInterface(version),
	}

	// read our input
	if err := h.input(input); err != nil {
		log.Error("failed to read input: %v", err)
	}

	// call the resource check function
	v, m, err := h.in.In(os.Args[1])
	if err != nil {
		log.Error("%v", err)
		return
	}

	result := inOutput{
		Metadata: m,
		Version:  v,
	}

	if err := h.output(result); err != nil {
		log.Error("failed to write response to concourse: %v", err)
	}
}

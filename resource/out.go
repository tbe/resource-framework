package resource

import (
	"os"

	ji "github.com/tbe/resource-framework/internal/jsoninterface"
	"github.com/tbe/resource-framework/log"
)

type outInput struct {
	Source *ji.Interface `json:"source"`
	Params *ji.Interface `json:"params"`
}

type outOutput struct {
	Metadata []interface{} `json:"metadata"`
	Version  interface{}   `json:"version"`
}

// Out calls the Out function of the resource implementation and handles the communication with concourse
func (h *Handler) Out() {
	// check that we are invoked with an commandline argument
	if len(os.Args) < 2 {
		log.Error("missing commandline argument")
		return
	}

	// check for a valid in resource
	if h.out == nil {
		_ = h.output(struct{}{})
		return
	}

	// get the storage for the source
	source := h.out.Source()
	if err := validateStructPtr(source); err != nil {
		log.Error("invalid source storage: %v", err)
	}

	// get the storage for the params
	params := h.out.Params()
	if err := validateStructPtr(params); err != nil {
		log.Error("invalid version storage: %v", err)
	}

	input := &outInput{
		Source: ji.NewInterface(source),
		Params: ji.NewInterface(params),
	}

	// read our input
	if err := h.input(input); err != nil {
		log.Error("failed to read input: %v", err)
	}

	// call the resource check function
	v, m, err := h.out.Out(os.Args[1])
	if err != nil {
		log.Error("%v", err)
		return
	}

	result := outOutput{
		Metadata: m,
		Version:  v,
	}

	if err := h.output(result); err != nil {
		log.Error("failed to write response to concourse: %v", err)
	}
}

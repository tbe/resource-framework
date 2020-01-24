package resource

import (
	"errors"
	"os"

	ji "github.com/tbe/resource-framework/internal/jsoninterface"
	"github.com/tbe/resource-framework/log"
)

type inInput struct {
	Source  *ji.Interface `json:"source"`
	Version *ji.Interface `json:"version"`
	Params  *ji.Interface `json:"params"`
}

type inOutput struct {
	Metadata []interface{} `json:"metadata"`
	Version  interface{}   `json:"version"`
}

// In calls the In function of the resource implementation and handles the communication with concourse
func (h *Handler) In() error {
	if len(os.Args) < 2 {
		err := errors.New("missing commandline argument")
		log.Error("%v", err)
		return err
	}

	// check for a valid in resource
	if h.in == nil {
		_ = h.output(struct{}{})
		return nil
	}

	// get the storage for the source
	source := h.in.Source()
	if err := validateStructPtr(source); err != nil {
		log.Error("invalid source storage: %v", err)
		return err
	}

	// get the storage for the version
	version := h.in.Version()
	if err := validateStructPtr(version); err != nil {
		log.Error("invalid version storage: %v", err)
		return err
	}

	// get the storage for the params
	params := h.in.Params()
	if err := validateStructPtr(params); err != nil {
		log.Error("invalid params storage: %v", err)
		return err
	}

	input := &inInput{
		Source:  ji.NewInterface(source),
		Version: ji.NewInterface(version),
		Params:  ji.NewInterface(params),
	}

	// read our input
	if err := h.input(input); err != nil {
		return err
	}

	// call the resource check function
	v, m, err := h.in.In(os.Args[1])
	if err != nil {
		log.Error(err.Error())
		return err
	}

	result := inOutput{
		Metadata: m,
		Version:  v,
	}

	if err := h.output(result); err != nil {
		return err
	}
	return nil
}

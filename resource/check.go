package resource

import (
	ji "github.com/tbe/resource-framework/internal/jsoninterface"
	"github.com/tbe/resource-framework/log"
)

type checkInput struct {
	Source  *ji.Interface `json:"source"`
	Version *ji.Interface `json:"version"`
}

type checkOutput []interface{}

// Check calls the Check function of the resource implementation and handles the communication with concourse
func (h *Handler) Check() error {
	// check if we have a valid resource for check
	if h.check == nil {
		_ = h.output(checkOutput{})
		return nil
	}
	// get the storage for the source
	source := h.check.Source()
	if err := validateStructPtr(source); err != nil {
		log.Error("invalid source storage: %v", err)
		return err
	}

	// get the storage for the version
	version := h.check.Version()
	if err := validateStructPtr(version); err != nil {
		log.Error("invalid version storage: %v", err)
		return err
	}

	input := &checkInput{
		Source:  ji.NewInterface(source),
		Version: ji.NewInterface(version),
	}

	// read our input
	if err := h.input(input); err != nil {
		return err
	}

	// call the resource check function
	result, err := h.check.Check()
	if err != nil {
		return err
	}

	if err := h.output(result); err != nil {
		return err
	}
	return nil
}

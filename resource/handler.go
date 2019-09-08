package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"gopkg.in/go-playground/validator.v9"

	ji "github.com/tbe/resource-framework/internal/jsoninterface"
	"github.com/tbe/resource-framework/log"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterStructValidation(ji.InterfaceValidator, ji.Interface{})
}

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
func (h *Handler) Run() error {
	switch path.Base(os.Args[0]) {
	case "check":
		return h.Check()
	case "in":
		return h.In()
	case "out":
		return h.Out()
	default:
		err := fmt.Errorf("unknown action %v", os.Args[0])
		log.Error(err.Error())
		return err
	}
}

func (h *Handler) output(data interface{}) error {
	if err := json.NewEncoder(h.Stdout).Encode(data); err != nil {
		err := fmt.Errorf("failed to write response to concourse: %v", err)
		log.Error(err.Error())
		return err
	}
	return nil
}

func (h *Handler) input(data interface{}) error {
	if err := json.NewDecoder(h.Stdin).Decode(data); err != nil {
		err := fmt.Errorf("failed to read input: %v", err)
		log.Error(err.Error())
		return err
	}

	// validate the struct
	if err := validate.Struct(data); err != nil {
		err := fmt.Errorf("failed to validate input: %v", err)
		log.Error(err.Error())
		return err
	}
	return nil
}

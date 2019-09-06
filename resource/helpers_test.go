package resource

import (
	"testing"
)

func TestValidateSource(t *testing.T) {
	// check for nil handling
	if err := validateStructPtr(nil); err == nil || err.Error() != "source is nil" {
		t.Error("failed to detect nil input")
	}

	// check for indirect nil handling
	if err := validateStructPtr((*struct{})(nil)); err == nil || err.Error() != "source is not a ptr to struct" {
		t.Error("failed to detect indirect nil input")
	}

	// check valid input
	if err := validateStructPtr(&struct{}{}); err != nil {
		t.Error("failed to validate struct", err)
	}
}

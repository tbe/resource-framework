package resource

import (
	"errors"
	"reflect"
)

func validateStructPtr(source interface{}) error {
	if source == nil {
		return errors.New("source is nil")
	}

	if reflect.TypeOf(source).Kind() != reflect.Ptr {
		return errors.New("source is not a ptr")
	}

	v := reflect.Indirect(reflect.ValueOf(source))

	if v.Kind() != reflect.Struct {
		return errors.New("source is not a ptr to struct")
	}

	return nil
}


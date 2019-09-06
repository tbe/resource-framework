package jsoninterface

import (
    "encoding/json"
    "reflect"
)

type Interface struct {
    data interface{}
}

func NewInterface(i interface{}) *Interface {
    return &Interface{data: i}
}

func (i *Interface) Data() interface{} {
    return i.data
}

func (i *Interface) UnmarshalJSON(data []byte) error {
    // get the reflect value and interface
    v := reflect.ValueOf(i.data).Interface()
    err := json.Unmarshal(data, v)
    if err != nil {
        return err
    }
    i.data = v
    return nil
}

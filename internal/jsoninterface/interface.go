package jsoninterface

import (
    "encoding/json"
    "reflect"

    "gopkg.in/go-playground/validator.v9"
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


func InterfaceValidator(sl validator.StructLevel) {
    jinterface := sl.Current().Interface().(Interface)
    if errs := sl.Validator().Struct(jinterface.data); errs != nil {
        for _,err := range errs.(validator.ValidationErrors) {
            sl.ReportError(err.Value(),err.Field(),err.StructField(),err.Tag(),err.Param())
        }
    }
}
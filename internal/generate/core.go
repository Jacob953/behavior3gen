package generate

import (
	"fmt"
	"github.com/magicsea/behavior3go/core"
	"reflect"
	"unsafe"
)

func SetBaseNodePrivate(node *core.BaseNode, desc, category string) (err error) {
	val := reflect.ValueOf(node).Elem()
	err = setField(val, "description", desc)
	if err != nil {
		return err
	}
	err = setField(val, "category", category)
	if err != nil {
		return err
	}
	return nil
}

func setField(v reflect.Value, name, val string) error {
	field := v.FieldByName(name)
	if !field.IsValid() {
		return fmt.Errorf("cannot find field %s, check object", name)
	}
	if field.CanSet() {
		field.SetString(val)
	} else {
		ptr := unsafe.Pointer(field.UnsafeAddr())
		reflect.NewAt(field.Type(), ptr).Elem().SetString(val)
	}
	return nil
}

package common

import (
	"fmt"
	"reflect"
)

// CopyStruct 構造体を別の構造体にコピー
func CopyStruct(src interface{}, dst interface{}) error {
	fv := reflect.ValueOf(src)

	ft := fv.Type()
	if fv.Kind() == reflect.Ptr {
		ft = ft.Elem()
		fv = fv.Elem()
	}

	tv := reflect.ValueOf(dst)
	if tv.Kind() != reflect.Ptr {
		return fmt.Errorf("[Error] non-pointer: %v", dst)
	}

	num := ft.NumField()
	for i := 0; i < num; i++ {
		field := ft.Field(i)

		if !field.Anonymous {
			name := field.Name
			srcField := fv.FieldByName(name)
			dstField := tv.Elem().FieldByName(name)

			if srcField.IsValid() && dstField.IsValid() {
				if srcField.Type() == dstField.Type() {
					dstField.Set(srcField)
				}
			}
		}
	}

	return nil
}

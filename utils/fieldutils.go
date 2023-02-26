package utils

import (
	"infra/types"
	"reflect"
)

func GetFieldColumnMap[T any](proto interface{}) map[string]interface{} {
	mapping := make(map[string]string)
	var t T
	mt := reflect.TypeOf(t)
	for i := 0; i < mt.NumField(); i++ {
		f := mt.Field(i)
		val := ParseTagSetting(f.Tag.Get(types.FieldTagType.Gorm), ";").GetSettingValue("column", "")
		if len(val) > 0 {
			mapping[f.Name] = val
		}
	}

	res := make(map[string]interface{})
	protoV := reflect.Indirect(reflect.ValueOf(proto))
	protoT := protoV.Type()
	for i := 0; i < protoV.NumField(); i++ {
		fv := protoV.Field(i)
		if fv.IsNil() {
			continue
		}

		coloumName := mapping[protoT.Field(i).Name]
		if len(coloumName) > 0 {
			res[coloumName] = fv.Elem().Interface()
		}
	}

	return res
}

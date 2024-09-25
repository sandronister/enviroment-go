package load

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func convertType(value string, field reflect.Value) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("cannot parse %s as int: %v", value, err)
		}
		field.SetInt(intValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("cannot parse %s as bool: %v", value, err)
		}
		field.SetBool(boolValue)
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("cannot parse %s as float: %v", value, err)
		}
		field.SetFloat(floatValue)
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}

	return nil
}

func (r *environment) getFieldName(field *reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return field.Name
	}

	return strings.Split(jsonTag, ",")[0]
}

func (r *environment) LoadVariable(result interface{}) error {
	v := reflect.ValueOf(result).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		jsonTag := r.getFieldName(&field)

		value, ok := r.variables[jsonTag]
		if !ok {
			return fmt.Errorf("variable %s not found", jsonTag)
		}

		fieldVal := v.FieldByName(field.Name)
		if !fieldVal.IsValid() {
			return fmt.Errorf("field %s not found in struct", field.Name)
		}

		if !fieldVal.CanSet() {
			return fmt.Errorf("cannot set field %s", field.Name)
		}

		if err := convertType(value, fieldVal); err != nil {
			return err
		}
	}

	return nil
}

package bengoder

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Encode(input interface{}) (string, error) {
	var result string

	err := writeFieldValue(&result, reflect.ValueOf(input))
	if err != nil {
		return "", err
	}

	return result, nil
}

func writeFieldValue(result *string, v reflect.Value) error {
	if !v.IsValid() {
		return errors.New("value is invalid")
	}

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		*result += "i"
		writeInt(result, v)
		*result += "e"
	case reflect.Struct:
		*result += "d"

		err := writeStruct(result, v)
		if err != nil {
			return err
		}

		*result += "e"
	case reflect.String:
		writeString(result, v)
	case reflect.Slice:
		*result += "l"
		writeSlice(result, v)
		*result += "e"
	default:
		return errors.New("unsupported value type " + v.Kind().String())
	}

	return nil
}

func writeInt(result *string, v reflect.Value) {
	fieldValue := v.Interface()

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		castedValue := fieldValue.(int)
		*result += strconv.Itoa(castedValue)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		castedValue := fieldValue.(uint)
		*result += strconv.FormatUint(uint64(castedValue), 10)
	case reflect.Int64:
		castedValue := fieldValue.(int64)
		*result += strconv.FormatInt(castedValue, 10)
	case reflect.Uint64:
		castedValue := fieldValue.(uint64)
		*result += strconv.FormatUint(castedValue, 10)
	}
}

func writeStruct(result *string, v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		tagName := v.Type().Field(i).Tag.Get("bengoder")

		if tagName != "" {
			fieldName = tagName
		}

		key := getStringLength(fieldName)
		value := fieldName

		writeKeyAndValue(result, key, value)
		err := writeFieldValue(result, v.Field(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func writeString(result *string, v reflect.Value) {
	*result += fmt.Sprintf("%s:%s", getStringLength(v.String()), v.String())
}

func writeSlice(result *string, v reflect.Value) {
	for i := 0; i < v.Len(); i++ {
		*result += "l"

		fieldName := v.Index(i).String()
		key := getStringLength(fieldName)
		value := fieldName

		writeKeyAndValue(result, key, value)

		*result += "e"
	}
}

func writeKeyAndValue(result *string, key string, value string) {
	*result += key + ":" + strings.ToLower(value)
}

func getStringLength(value string) string {
	return strconv.Itoa(len(value))
}

package strunpack

import (
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Unpack unpacks a string into a struct using a regular expression.
func Unpack(s string, re *regexp.Regexp, res interface{}) error {
	if re == nil {
		return errors.Errorf("non-nil regexp required")
	}

	ms := re.FindStringSubmatch(s)
	if len(ms) < 2 {
		return errors.Errorf("No matches found")
	}

	resType := reflect.TypeOf(res)
	if resType.Kind() != reflect.Ptr || resType.Elem().Kind() != reflect.Struct {
		return errors.Errorf("Invalid result type. Expected a pointer to a struct")
	}
	resValue := reflect.ValueOf(res).Elem()

	subExpNames := re.SubexpNames()
	numSubExpNames := 0
	for _, s := range subExpNames {
		if strings.TrimSpace(s) != "" {
			numSubExpNames++
		}
	}
	{
		log.Printf("subExpNames: %d %v : %d", numSubExpNames, subExpNames, len(subExpNames))
	}
	if numSubExpNames == len(ms)-1 {
		for i, name := range re.SubexpNames()[1:] {
			fieldName := cases.Title(language.English).String(name)
			field := resValue.FieldByName(fieldName)
			if !field.IsValid() || !field.CanSet() {
				continue
			}
			typ := field.Type()
			fieldStruct, ok := resType.Elem().FieldByName(fieldName)
			if !ok {
				return errors.Errorf("Internal error: Could not find field %s in struct", fieldName)
			}
			m := ms[i+1]
			val, err := valueOf(m, fieldStruct)
			if err != nil {
				return err
			}
			if !val.Type().ConvertibleTo(typ) {
				return errors.Errorf("Incompatible types between struct field(%s) and matched value(%v)", name, val)
			}
			field.Set(val.Convert(typ))
		}
	} else if numSubExpNames == 0 {
		for i, m := range ms[1:] {
			field := resValue.FieldByIndex([]int{i})
			if !field.IsValid() || !field.CanSet() {
				continue
			}
			typ := field.Type()
			fieldStruct := resType.Elem().FieldByIndex([]int{i})
			val, err := valueOf(m, fieldStruct)
			if err != nil {
				return err
			}
			if !val.Type().ConvertibleTo(typ) {
				return errors.Errorf("Incompatible types between struct field(%d:%v) and matched value(%v)", i, field, val)
			}
			field.Set(val.Convert(typ))
		}
	} else {
		return errors.Errorf("Internal error: Number of subexp names (%d) must either match number of matches (%d) or be 0",
			len(subExpNames), len(ms))
	}

	return nil
}

func valueOf(s string, field reflect.StructField) (reflect.Value, error) {
	switch field.Type.Kind() {
	case reflect.String:
		return reflect.ValueOf(s), nil
	case reflect.Int:
		i, err := strconv.Atoi(s)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(i), nil
	case reflect.Int8:
		i, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(int8(i)), nil
	case reflect.Int16:
		i, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(int16(i)), nil
	case reflect.Int32:
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(int32(i)), nil
	case reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(i), nil
	case reflect.Float32:
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(float32(f)), nil
	case reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(f), nil
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(b), nil
	default:
		return reflect.Value{}, errors.Errorf("Unsupported type: %s", field.Type.Kind())
	}
}

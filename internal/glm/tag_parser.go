package glm

import (
	"fmt"
	"reflect"
	"strings"
)

// parseGLMTags extracts glm tag values from a struct and returns a map of them
func parseGLMTags(input interface{}) ([]string, error) {
	result := make([]string, 0)
	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// check if input is a struct
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct or pointer to a struct")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tagValue, ok := field.Tag.Lookup("glm")
		if !ok {
			continue
		}

		hasOmit, _ := hasOption(tagValue, "omitempty")
		empty := isEmpty(val.Field(i))
		hasDefault, defaultValue := hasOption(tagValue, "default")

		// if "omitempty" is present and the field value is zero, skip the field
		if hasOmit && empty {
			continue
		}

		value := stringify(val.Field(i))

		if hasDefault && empty {
			value = defaultValue
		}

		// get the main value of the tag (before the comma)
		name := getName(tagValue)
		if name != "" {
			result = append(result, fmt.Sprintf("%s %s", name, value))
		}
	}

	return result, nil
}

// stringify converts a reflect.Value to its string representation.
func stringify(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v.Float())
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}

// Return true if the tag is present, and a string value if there is an equal
func hasOption(tag string, option string) (bool, string) {
	for _, opt := range strings.Split(tag, ",")[1:] {
		// find the value after an equal sign if the equal is present
		equalSplit := strings.Split(opt, "=")
		value := ""
		if len(equalSplit) > 1 {
			opt = equalSplit[0]
			value = equalSplit[1]
		}

		if opt == option {
			return true, value
		}
	}
	return false, ""
}

func getName(tag string) string {
	return strings.Split(tag, ",")[0]
}

func isEmpty(value reflect.Value) bool {
	return value.Interface() == reflect.Zero(value.Type()).Interface()
}

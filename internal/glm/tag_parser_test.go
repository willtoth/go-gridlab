package glm

import (
	"reflect"
	"testing"
)

func TestParseGLMTags(t *testing.T) {
	type MyStruct struct {
		Name             string `glm:"name"`
		Age              int    `glm:"age,omitempty"`
		Empty            string `glm:"empty,omitempty"`
		NotInTheMap      string
		SomeFloat        float64 `glm:"some_float,omitempty"`
		SomeBool         bool    `glm:"some_bool,omitempty"`
		HasDefault       string  `glm:"my_value,default=foo"`
		HasUnusedDefault string  `glm:"my_value2,default=bar"`
		EmptyFloat       float64 `glm:"empty_float,omitempty"`
	}

	person := MyStruct{
		Name:             "John",
		Age:              0,
		SomeFloat:        1.25,
		SomeBool:         true,
		NotInTheMap:      "not in the map",
		HasUnusedDefault: "here",
	}

	result, err := parseGLMTags(person)
	if err != nil {
		t.Fatalf("Failed to parse GLM tags: %v", err)
	}

	expected := []string{
		"name John",
		"some_float 1.250000",
		"some_bool true",
		"my_value foo",
		"my_value2 here",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}
}

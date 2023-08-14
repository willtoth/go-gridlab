package glm

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildWithOptions(t *testing.T) {
	builder := NewBlockBuilder("testBlock", ObjectType)
	builder.AddOption("option1", "value1").AddOption("option2", "value2")

	expected := `object testBlock {
	option1 value1
	option2 value2
}`

	output := builder.Build()
	if diff := cmp.Diff(expected, output); diff != "" {
		t.Fatalf("Mismatch (-want +got):\n%s", diff)
	}
}

func TestBuildWithNoOptions(t *testing.T) {
	builder := NewBlockBuilder("testBlock", ObjectType)

	expected := "object testBlock;"
	output := builder.Build()
	if output != expected {
		t.Errorf("Expected output '%s', but got '%s'", expected, output)
	}
}

func TestBuildWithStruct(t *testing.T) {
	builder := NewBlockBuilder("testBlock", ObjectType)
	builder.AddStruct(&struct {
		Name        string `glm:"name"`
		Age         int    `glm:"age,omitempty"`
		Empty       string `glm:"empty,omitempty"`
		NotInTheMap string
		SomeFloat   float64 `glm:"some_float,omitempty"`
		SomeBool    bool    `glm:"some_bool,omitempty"`
	}{
		Name:        "John",
		Age:         0,
		SomeFloat:   1.25,
		SomeBool:    true,
		NotInTheMap: "not in the map",
	})

	expected := `object testBlock {
	name John
	some_float 1.250000
	some_bool true
}`
	output := builder.Build()
	if diff := cmp.Diff(expected, output); diff != "" {
		t.Fatalf("Mismatch (-want +got):\n%s", diff)
	}
}

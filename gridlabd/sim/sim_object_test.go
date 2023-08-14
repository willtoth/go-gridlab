package sim

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBasicGLMSim(t *testing.T) {
	// create a template class

	climate := &Climate{
		Name:    "climate",
		Tmyfile: "tmyfile",
	}

	house := &House{
		Name:            "house",
		Parent:          "parent",
		HvacPowerFactor: 3.90,
		FloorAreaSqFt:   1000,
		CoolingSetpoint: 70,
	}

	expected := `object house {
	name house
	parent parent
	hvac_power_factor 3.900000
	floor_area 1000
	cooling_setpoint 70.000000
}
object climate {
	name climate
	tmyfile tmyfile
	interpolate LINEAR
}
`

	output := GenerateGlm([]SimObject{house, climate})
	if diff := cmp.Diff(expected, output); diff != "" {
		t.Fatalf("Mismatch (-want +got):\n%s", diff)
	}
}

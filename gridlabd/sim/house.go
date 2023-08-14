package sim

type House struct {
	Name            string  `glm:"name"`
	Parent          string  `glm:"parent"`
	HvacPowerFactor float64 `glm:"hvac_power_factor,omitempty"`
	FloorAreaSqFt   int32   `glm:"floor_area,omitempty"`
	CoolingSetpoint float64 `glm:"cooling_setpoint,omitempty"`
	HeatingSetpoint float64 `glm:"heating_setpoint,omitempty"`
	AirTemperature  float64 `glm:"air_temperature,omitempty"`
}

func (House) ClassName() string {
	return "house"
}

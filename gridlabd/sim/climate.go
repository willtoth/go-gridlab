package sim

type Climate struct {
	Name        string `glm:"name"`
	Tmyfile     string `glm:"tmyfile"`
	Interpolate string `glm:"interpolate,default=LINEAR"`
}

func (climate *Climate) ClassName() string {
	return "climate"
}

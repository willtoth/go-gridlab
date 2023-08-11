package gridlabd

type ClimateData struct {
	ID                               int     `json:"id"`
	Class                            string  `json:"class"`
	Name                             string  `json:"name"`
	Rank                             int     `json:"rank"`
	Clock                            int     `json:"clock"`
	ValidTo                          int     `json:"valid_to"`
	Latitude                         float64 `json:"latitude"`
	Longitude                        float64 `json:"longitude"`
	Flags                            string  `json:"flags"`
	SolarElevation                   string  `json:"solar_elevation"`
	SolarAzimuth                     string  `json:"solar_azimuth"`
	SolarZenith                      string  `json:"solar_zenith"`
	City                             string  `json:"city"`
	Tmyfile                          string  `json:"tmyfile"`
	Temperature                      string  `json:"temperature"`
	Humidity                         string  `json:"humidity"`
	SolarFlux                        string  `json:"solar_flux"`
	SolarDirect                      string  `json:"solar_direct"`
	SolarDiffuse                     string  `json:"solar_diffuse"`
	SolarGlobal                      string  `json:"solar_global"`
	ExtraterrestrialGlobalHorizontal string  `json:"extraterrestrial_global_horizontal"`
	ExtraterrestrialDirectNormal     string  `json:"extraterrestrial_direct_normal"`
	Pressure                         string  `json:"pressure"`
	WindSpeed                        string  `json:"wind_speed"`
	WindDir                          string  `json:"wind_dir"`
	WindGust                         string  `json:"wind_gust"`
	RecordLow                        string  `json:"record_low"`
	RecordLowDay                     int     `json:"record_low_day"`
	RecordHigh                       string  `json:"record_high"`
	RecordHighDay                    int     `json:"record_high_day"`
	RecordSolar                      string  `json:"record_solar"`
	Rainfall                         string  `json:"rainfall"`
	Snowdepth                        string  `json:"snowdepth"`
	Interpolate                      string  `json:"interpolate"`
	SolarHoriz                       string  `json:"solar_horiz"`
	SolarNorth                       string  `json:"solar_north"`
	SolarNortheast                   string  `json:"solar_northeast"`
	SolarEast                        string  `json:"solar_east"`
	SolarSoutheast                   string  `json:"solar_southeast"`
	SolarSouth                       string  `json:"solar_south"`
	SolarSouthwest                   string  `json:"solar_southwest"`
	SolarWest                        string  `json:"solar_west"`
	SolarNorthwest                   string  `json:"solar_northwest"`
	SolarRaw                         string  `json:"solar_raw"`
	GroundReflectivity               string  `json:"ground_reflectivity"`
	Reader                           string  `json:"reader"`
	Forecast                         string  `json:"forecast"`
	CloudModel                       string  `json:"cloud_model"`
	CloudOpacity                     string  `json:"cloud_opacity"`
	OpqSkyCov                        string  `json:"opq_sky_cov"`
	CloudSpeedFactor                 string  `json:"cloud_speed_factor"`
	SolarCloudDirect                 string  `json:"solar_cloud_direct"`
	SolarCloudDiffuse                string  `json:"solar_cloud_diffuse"`
	SolarCloudGlobal                 string  `json:"solar_cloud_global"`
	CloudAlpha                       string  `json:"cloud_alpha"`
	CloudNumLayers                   int     `json:"cloud_num_layers"`
	CloudAerosolTransmissivity       string  `json:"cloud_aerosol_transmissivity"`
	HeatIndex                        string  `json:"heat_index"`
}

func (cd ClimateData) className() string {
	return "climate"
}

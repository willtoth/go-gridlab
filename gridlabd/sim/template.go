package sim

import (
	"bytes"
	"text/template"
)

type SimTemplate struct {
}

func (st *SimTemplate) ClassName() string {
	return "SimTemplate"
}

func (st *SimTemplate) generateGlm() (string, error) {
	tmpl, err := template.New("simTemplate").Parse(glmBaseTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, st)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Basic template to get things moving, assume essentially perfect generation
// and generate triplex meter data + residential
const glmBaseTemplate = `
module powerflow {
	solver_method NR;
	NR_iteration_limit 50;
};

module tape;
module climate;
module residential;

clock {
	timezone PST+8PDT;
	starttime '2009-08-08 04:30:00';
	stoptime '2009-08-09 04:30:00';
}

//Configurations
object line_configuration {
	name line_configuration_1;
	conductor_A overhead_line_conductor_1;
	conductor_B overhead_line_conductor_1;
	conductor_C overhead_line_conductor_1;
	conductor_N overhead_line_conductor_1;
	spacing line_spacing_1;
}
object line_spacing {
	name line_spacing_1;
	distance_AB 57.6 in;
	distance_BC 57.6 in;
	distance_AC 57.6 in;
	distance_AN 51.6 in;
	distance_BN 51.6 in;
	distance_CN 51.6 in;
}
object overhead_line_conductor {
	name overhead_line_conductor_1;
	rating.summer.continuous 256.0;
	geometric_mean_radius 0.01200 ft;
	resistance 0.1;
}
object transformer_configuration {
	name substation_transformer_config;
	connect_type WYE_WYE;
	install_type PADMOUNT;
	primary_voltage 132790;
	secondary_voltage 7216;
	power_rating 8.4 MVA;
	impedance 0.000033+0.0022j;
}
object transformer_configuration {
	name house_transformer;
	connect_type SINGLE_PHASE_CENTER_TAPPED;
	install_type PADMOUNT;
	primary_voltage 7200 V;
	secondary_voltage 230 V;
	power_rating 250;
	powerA_rating 250;
	impedance 0.0015+0.0675j;
}
object triplex_line_configuration {
	name triplex_line_config;
	conductor_1 Name_1_0_AA_triplex;
	conductor_2 Name_1_0_AA_triplex;
	conductor_N Name_1_0_AA_triplex;
	insulation_thickness 0.1;
	diameter 0.4;
}
object triplex_line_conductor {
	name Name_1_0_AA_triplex;
	resistance 0.57;
	geometric_mean_radius 0.0111;
}


//Feeder head
object node {
	name feeder_head;
	bustype SWING;
	phases ABCN;
	nominal_voltage 132790;
}

object transformer {
	name substation_transformer;
	from feeder_head;
	to substation_meter;
	phases ABCN;
	configuration substation_transformer_config;
}

object meter {
	name substation_meter;
	phases ABCN;
	nominal_voltage 7216.88;
}

//Branch
object overhead_line {
	name branch_1_line_1;
	phases ABCN;
	from substation_meter;
	to branch_1_meter_1;
	length 1000;
	configuration line_configuration_1;
}
object meter {
	name branch_1_meter_1;
	groupid branch_1_meter;
	phases ABCN;
	nominal_voltage 7216.88;
}

// Branch 1 house 1
object transformer {
	name b1m1_house_trans;
	phases AS;
	from branch_1_meter_1;
	to b1m1_house_node;
	configuration house_transformer;
}
object triplex_node {
	name b1m1_house_node;
	phases AS;
	nominal_voltage 124.00;
}
object triplex_line {
	name b1m1_tl;
	from b1m1_house_node;
	to b1m1_house_meter;
	phases AS;
	length 10;
	configuration triplex_line_config;
}
object triplex_meter {
	name b1m1_house_meter;
	phases AS;
	nominal_voltage 230.00;
}
object house {
	name b1m1_house;
	parent b1m1_house_meter;
	thermal_integrity_level LITTLE;
	hvac_power_factor 0.97;
	cooling_COP 3.90;
	floor_area 1040;
	cooling_setpoint 75;
	thermostat_deadband 2;
	air_temperature 72.5;
}

`

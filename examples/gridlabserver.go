package main

import (
	"fmt"

	"github.com/willtoth/go-gridlab/gridlabd"
)

func main() {
	gld := &gridlabd.Gridlabd{
		Hostname: "localhost",
		Port:     6267,
	}

	some, _ := gridlabd.Find[gridlabd.TriplexMeterData](gld)

	for _, a := range some {
		fmt.Println("Bill", a.MonthlyBill)
	}
}

package main

import (
	"example.com/gotax/prices"
)

func main() {

	// Indian GST rates
	gstRates := []float64{0.00, 0.05, 0.12, 0.18, 0.28}

	for _, gstRate := range gstRates {
		calculation := prices.New(gstRate)
		calculation.Process()

	}

}

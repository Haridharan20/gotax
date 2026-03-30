package main

import (
	"fmt"

	"example.com/gotax/filemanager"
	"example.com/gotax/prices"
)

func main() {
	// Indian GST rates
	gstRates := []float64{0.00, 0.05, 0.12, 0.18, 0.28}

	for _, gstRate := range gstRates {
		fm := filemanager.FileManager{
			InputFile:  "prices.txt",
			OutputFile: fmt.Sprintf("result_%.0f.json", gstRate*100),
		}

		calculation := prices.New(gstRate, fm)
		calculation.Process()
	}

	// Example: Use CMDManager for command-line input
	// cm := cmdmanager.CMDManager{}
	// for _, gstRate := range gstRates {
	// 	calculation := prices.New(gstRate, cm)
	// 	calculation.Process()
	// }
}

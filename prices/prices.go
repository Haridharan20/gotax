package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var gstCategories = map[float64]string{
	0.00: "Essentials (Food, Medicine)",
	0.05: "Basic Items",
	0.12: "Most Items",
	0.18: "Services & Standard Items",
	0.28: "Luxury Items",
}

type GSTCalculation struct {
	GSTRate       float64
	Category      string
	BasePrices    []float64
	PricesWithGST map[string]string
}

func (calc *GSTCalculation) LoadPricesFromFile() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIdx, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Converting price to float failed.")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIdx] = floatPrice
	}

	calc.BasePrices = prices
}

func (calc *GSTCalculation) Process() {
	calc.LoadPricesFromFile()
	calc.PricesWithGST = make(map[string]string)

	for _, price := range calc.BasePrices {
		gstIncludedPrice := price * (1 + calc.GSTRate)
		calc.PricesWithGST[fmt.Sprintf("₹%.2f", price)] = fmt.Sprintf("₹%.2f", gstIncludedPrice)
	}

	fmt.Printf("Category: %s | GST (%.0f%%): %v\n", calc.Category, calc.GSTRate*100, calc.PricesWithGST)
}

func New(gstRate float64) *GSTCalculation {
	return &GSTCalculation{
		BasePrices: []float64{100, 500, 1000},
		GSTRate:    gstRate,
		Category:   gstCategories[gstRate],
	}
}

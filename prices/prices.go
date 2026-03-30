package prices

import (
	"fmt"

	"example.com/gotax/utils"
)

var gstCategories = map[float64]string{
	0.00: "Essentials (Food, Medicine)",
	0.05: "Basic Items",
	0.12: "Most Items",
	0.18: "Services & Standard Items",
	0.28: "Luxury Items",
}

type GSTCalculation struct {
	GSTRate       float64            `json:"gst_rate"`
	Category      string             `json:"category"`
	BasePrices    []float64          `json:"base_prices"`
	PricesWithGST map[string]string  `json:"prices_with_gst"`
}

func (calc *GSTCalculation) LoadPricesFromFile() {

	lines, err := utils.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := utils.StringToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	calc.BasePrices = prices
}

func (calc *GSTCalculation) Process() {
	calc.LoadPricesFromFile()
	result := make(map[string]string)

	for _, price := range calc.BasePrices {
		gstIncludedPrice := price * (1 + calc.GSTRate)
		result[fmt.Sprintf("₹%.2f", price)] = fmt.Sprintf("₹%.2f", gstIncludedPrice)
	}

	calc.PricesWithGST = result

	utils.WriteJson(fmt.Sprintf("result_%.0f.json", calc.GSTRate*100), calc)

	// fmt.Printf("Category: %s | GST (%.0f%%): %v\n", calc.Category, calc.GSTRate*100, calc.PricesWithGST)
}

func New(gstRate float64) *GSTCalculation {
	return &GSTCalculation{
		BasePrices: []float64{100, 500, 1000},
		GSTRate:    gstRate,
		Category:   gstCategories[gstRate],
	}
}

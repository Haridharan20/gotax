package prices

import "fmt"

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
	PricesWithGST map[string]float64
}

func (calc *GSTCalculation) Process() {
	calc.PricesWithGST = make(map[string]float64)

	for _, price := range calc.BasePrices {
		calc.PricesWithGST[fmt.Sprintf("₹%.2f", price)] = price * (1 + calc.GSTRate)
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

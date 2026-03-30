package prices

import (
	"fmt"
)

var gstCategories = map[float64]string{
	0.00: "Essentials (Food, Medicine)",
	0.05: "Basic Items",
	0.12: "Most Items",
	0.18: "Services & Standard Items",
	0.28: "Luxury Items",
}

type IOManager interface {
	LoadPrices() ([]float64, error)
	SaveResult(data any) error
}

type GSTCalculation struct {
	Manager       IOManager         `json:"-"`
	GSTRate       float64           `json:"gst_rate"`
	Category      string            `json:"category"`
	BasePrices    []float64         `json:"base_prices"`
	PricesWithGST map[string]string `json:"prices_with_gst"`
}

func (calc *GSTCalculation) Process() error {
	prices, err := calc.Manager.LoadPrices()
	if err != nil {
		return err
	}

	calc.BasePrices = prices
	result := make(map[string]string)

	for _, price := range calc.BasePrices {
		gstIncludedPrice := price * (1 + calc.GSTRate)
		result[fmt.Sprintf("₹%.2f", price)] = fmt.Sprintf("₹%.2f", gstIncludedPrice)
	}

	calc.PricesWithGST = result

	return calc.Manager.SaveResult(calc)
}

func New(gstRate float64, manager IOManager) *GSTCalculation {
	return &GSTCalculation{
		Manager:    manager,
		BasePrices: []float64{100, 500, 1000},
		GSTRate:    gstRate,
		Category:   gstCategories[gstRate],
	}
}

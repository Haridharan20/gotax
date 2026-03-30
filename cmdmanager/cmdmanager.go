package cmdmanager

import (
	"fmt"

	"example.com/gotax/utils"
)

type CMDManager struct {
}

func (cm CMDManager) LoadPrices() ([]float64, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return utils.StringToFloat(prices)
}

func (cm CMDManager) SaveResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}

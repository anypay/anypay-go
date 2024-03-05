// file: prices/prices.go
package prices

import "fmt"

type Price struct {
	ID       int
	Amount   float64
	Currency string
}

var allPrices = []Price{
	{ID: 1, Amount: 10.0, Currency: "USD"},
	{ID: 2, Amount: 15.0, Currency: "EUR"},
}

func ListAll() []Price {
	return allPrices
}

func FindByID(id string) (*Price, error) {
	for _, price := range allPrices {
		if fmt.Sprintf("%d", price.ID) == id {
			return &price, nil
		}
	}
	return nil, fmt.Errorf("price with ID %s not found", id)
}

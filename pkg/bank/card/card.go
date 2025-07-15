package card

import "github.com/olim007/bank/pkg/bank/types"

func Issue(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		Name:     name,
		Number:   "5555 xxxx xxxx 2222",
		Currency: currency,
		Color:    color,
		Balance:  0,
		Active:   true,
	}
}

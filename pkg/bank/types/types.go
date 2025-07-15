package types

type Money int64

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}

type Currency string

type PAN string

type Card struct {
	ID       int
	Name     string
	Number   PAN
	Currency Currency
	Color    string
	Balance  Money
	Active   bool
}

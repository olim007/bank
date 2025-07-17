package types

type Money int64

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}

const (
  TJS Currency = "TJS"
  RUB Currency = "RUB"
  USD Currency = "USD"
  EUR Currency = "EUR"
)

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

type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

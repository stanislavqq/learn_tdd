package currencyCalc

type Currency interface {
	GetAmount() int
	Equals(currency Currency) bool
	Times(amount int) Currency
	Currency() string
}

type Money struct {
	amount int
}

func NewMoney(amount int, currency string) Currency {
	switch currency {
	case "USD":
		return NewDollar(amount)
	case "CHF":
		return NewFranc(amount)
	}

	return nil
}

func NewDollar(amount int) Currency {
	return &Dollar{amount: amount, currency: "USD"}
}

func NewFranc(amount int) Currency {
	return &Franc{amount: amount, currency: "CHF"}
}

func (c *Money) GetAmount() int {
	return c.amount
}

package currencyCalc

type Franc struct {
	amount   int
	currency string
}

func (c *Franc) Times(count int) Currency {
	return NewFranc(5 * count)
}

func (c *Franc) Equals(currency Currency) bool {
	return c.amount == currency.GetAmount()
}

func (c *Franc) GetAmount() int {
	return c.amount
}

func (c *Franc) Currency() string {
	return "CHF"
}

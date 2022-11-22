package currencyCalc

type Franc struct {
}

func NewFranc(amount int) Money {
	franc := Franc{}
	return Money{Type: franc, amount: amount}
}

func (c *Franc) Times(count int) Money {
	return NewFranc(5 * count)
}

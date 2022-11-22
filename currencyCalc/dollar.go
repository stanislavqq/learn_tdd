package currencyCalc

import "reflect"

type Dollar struct {
	amount int
}

func NewDollar(amount int) Money {
	dollar := Dollar{amount: amount}
	return Money{Type: dollar, amount: amount}
}

func (d *Dollar) Times(count int) Money {
	return NewDollar(d.amount * count)
}

func (d *Dollar) Equals(currency Money) bool {
	return d.amount == currency.GetAmount() && reflect.TypeOf(currency) == reflect.TypeOf(d)
}

func (d *Dollar) GetAmount() int {
	return d.amount
}

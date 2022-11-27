package currencyCalc

import "reflect"

type Dollar struct {
	Money
	amount   int
	currency string
}

//func NewDollar(amount int) Currency {
//	return &Dollar{amount: amount}
//}

func (d *Dollar) Times(count int) Currency {
	return NewDollar(d.amount * count)
}

func (d *Dollar) Equals(currency Currency) bool {
	return d.amount == currency.GetAmount() && (reflect.TypeOf(d) == reflect.TypeOf(currency))
}

func (d *Dollar) GetAmount() int {
	return d.amount
}

func (d *Dollar) Currency() string {
	return "USD"
}

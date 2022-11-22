package currencyCalc

import "reflect"

//type Currency interface {
//	GetAmount() int
//	Equals(currency Currency) bool
//	Times(amount int) Currency
//}

type Money struct {
	Type   interface{}
	amount int
}

func (m *Money) Equals(money Money) bool {
	return m.amount == money.amount && reflect.TypeOf(money) == reflect.TypeOf(m)
}

func (m *Money) GetAmount() int {
	return m.amount
}

func (m *Money) Times(amount int) Money {
	return Money{Type: m.Type, amount: amount}
}

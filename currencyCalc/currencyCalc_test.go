package currencyCalc

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, NewFranc(10), five.Times(2))
	assert.Equal(t, NewFranc(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	t.Run("5 == 5 is true", func(t *testing.T) {
		fiveDollar := NewDollar(5)
		assert.True(t, fiveDollar.Equals(NewDollar(5)))

		fiveFranc := NewFranc(5)
		assert.True(t, fiveFranc.Equals(NewFranc(5)))
	})

	t.Run("5 == 6 is false", func(t *testing.T) {
		fiveDollar := NewDollar(5)
		assert.False(t, fiveDollar.Equals(NewDollar(6)))

		fiveFranc := NewFranc(5)
		assert.False(t, fiveFranc.Equals(NewFranc(6)))
	})

	t.Run("$5 == 5 CHF is false", func(t *testing.T) {
		fiveDollar := NewDollar(5)
		assert.False(t, fiveDollar.Equals(NewFranc(5)))
	})
}

func TestMoneyStructs(t *testing.T) {
	money := Money{amount: 5}
	franc := NewFranc(5)

	assert.True(t, reflect.TypeOf(money) == reflect.TypeOf(franc))
}

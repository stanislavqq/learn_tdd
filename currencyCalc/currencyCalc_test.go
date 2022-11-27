package currencyCalc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplication(t *testing.T) {
	assert.Equal(t, NewDollar(10), NewDollar(5).Times(2))
	assert.Equal(t, NewDollar(15), NewDollar(5).Times(3))
}

func TestFrancMultiplication(t *testing.T) {
	assert.Equal(t, NewFranc(10), NewFranc(5).Times(2))
	assert.Equal(t, NewFranc(15), NewFranc(5).Times(3))
}

func TestEquality(t *testing.T) {
	t.Run("5 == 5 is true", func(t *testing.T) {
		assert.True(t, NewDollar(5).Equals(NewDollar(5)))
		assert.True(t, NewFranc(5).Equals(NewFranc(5)))
	})

	t.Run("5 == 6 is false", func(t *testing.T) {
		assert.False(t, NewDollar(5).Equals(NewDollar(6)))
		assert.False(t, NewFranc(5).Equals(NewFranc(6)))
	})

	t.Run("$5 == 5 CHF is false", func(t *testing.T) {
		assert.False(t, NewDollar(5).Equals(NewFranc(5)))
	})
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency())
	assert.Equal(t, "CHF", NewFranc(1).Currency())
}

func TestDifferentClassEquality(t *testing.T) {
	assert.True(t, NewMoney(10, "CHF").Equals(NewFranc(10)))
}

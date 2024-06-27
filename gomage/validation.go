package gomage

import (
	"errors"
	"strconv"
)

// ValidateCardCvv validates CVV card code
func (c *Card) ValidateCardCvv() error {
	if len(c.Cvv) < 3 || len(c.Cvv) > 4 {
		return errors.New("invalid cvv length")
	}

	return nil
}

// ValidateCardNumber validates card numbers using luhn algorithm
func (c *Card) ValidateCardNumber() bool {
	var sum int
	var alt bool

	num := len(c.Number)
	if num < 13 || num > 19 {
		return false
	}

	for i := len(c.Number) - 1; i >= 0; i-- {
		mod, _ := strconv.Atoi(string(c.Number[i]))
		if alt {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alt = !alt

		sum += mod
	}

	return sum%10 == 0
}

package gomage

import (
	"errors"
	"fmt"
	"strconv"
)

// At returns the digits from the start to the given length
func (d *digits) At(i int) int {
	return d[i-1]
}

// Card holds the general information about credit card
type Card struct {
	Number, Cvv, Month, Year string
	Company
}

// Company holds a short and long names of who has issued the credit card
type Company struct {
	Short, Long string
}

// LastFourNumbers returns the last four number of credit card
func (c *Card) LastFourNumbers() (string, error) {
	if len(c.Number) < 4 {
		return "", errors.New("number too short")
	}

	fmt.Println(c.Number[len(c.Number)-4:])
	return c.Number[len(c.Number)-4:], nil
}

// New returns the credit card with false/nullified/generic information
func (c *Card) New() {
	c.Cvv, c.Month, c.Year, c.Number = "000", "01", "1970", "0000000000000000"
}

func (c *Card) Validate(allowNumbers ...bool) error {

	return nil
}

// ValidateExpiration validates the credit card's expiration date
func (c *Card) ValidateExpiration() error {
	var year, month int
	var err error
	timeNow := timeNowConfig()

	if len(c.Year) < 3 {
		year, err = strconv.Atoi(strconv.Itoa(timeNow.UTC().Year())[:2] + c.Year)
		if err != nil {
			return errors.New("invalid year")
		}
	} else {
		year, err = strconv.Atoi(c.Year)
		if err != nil {
			return errors.New("invalid year")
		}
	}

	month, err = strconv.Atoi(c.Month)
	if err != nil {
		return errors.New("invalid month")
	}

	if month < 1 || 12 < month {
		return errors.New("invalid month")
	}

	if year < timeNowConfig().UTC().Year() {
		return errors.New("credit card has expired")
	}

	if year == timeNowConfig().UTC().Year() && month < int(timeNowConfig().UTC().Month()) {
		return errors.New("credit card has expired")
	}

	return nil
}

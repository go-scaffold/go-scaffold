package prompt

import (
	"errors"
	"strconv"
)

func validateInteger(value string) error {
	_, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errors.New("Invalid number: " + value)
	}
	return nil
}

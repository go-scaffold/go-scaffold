package promptcli

import (
	"strconv"
)

func validateInteger(value string) error {
	_, err := strconv.ParseInt(value, 10, 64)
	return err
}

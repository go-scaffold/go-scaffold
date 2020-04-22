package promptcli

import (
	"github.com/manifoldco/promptui"
)

func validateFuncForType(promptType string) promptui.ValidateFunc {
	switch promptType {
	case "int":
		return validateInteger
	default:
		return nil
	}
}

package utils

import (
	"fmt"
)

func GetUserInputs(promptMsg string) string {
	fmt.Print(promptMsg)
	var value string
	fmt.Scanln(&value) //we do not need error from scan
	return value
}

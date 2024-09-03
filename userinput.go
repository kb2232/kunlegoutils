package utils

import (
	"errors"
	"fmt"
)

func GetUserInputs(promptMsg string) (string, error) {
	fmt.Print(promptMsg)
	var value string
	_, err := fmt.Scanln(&value)
	if err != nil {
		println(err)
		return "", errors.New("invalid input")
	}
	return value, nil
}

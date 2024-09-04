package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInputs(promptMsg string) string {
	fmt.Printf("%v ", promptMsg)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n") //remove endline
	text = strings.TrimSuffix(text, "\r") //remove endline in windows

	return text
}

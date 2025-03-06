package internal

import "fmt"

func ValidateArgs(command string, args []string, expected int) (bool, string) {
	if len(args) < expected {
		return false, fmt.Sprintf("ERROR: %s requires %d arguments", command, expected)
	}
	return true, ""
}

func IsEven(n int) bool {
	return n%2 == 0
}

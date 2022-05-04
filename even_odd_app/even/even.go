package even

import (
	"errors"
)

func Even(num int) (string, error) {
	var response string
	// if number is empty return empty message
	if num == 0 {
		return "", errors.New("0")
	}

	if num%2 == 0 {
		response = "Even"
	} else {
		response = "Odd"
	}
	return response, nil
}

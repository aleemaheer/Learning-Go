package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	response := fmt.Sprintf(randomFormat(), name)
	return response, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v",
		"Nice to meet you %v",
		"Happy to see you %v",
	}
	return formats[rand.Intn(len(formats))]
}

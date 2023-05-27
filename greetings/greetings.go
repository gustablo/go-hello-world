package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person. (this is documentation of method) (exported bc first letter is capitalized)
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func MultipleHellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages randomically. (not exported bc first letter is not capitalized)
func randomFormat() string {
	formats := []string{
		"Hi %v. Welcome!",
		"Great to see you %v!",
		"Hello %v",
	}

	randomic := rand.Intn(len(formats))

	return formats[randomic]
}

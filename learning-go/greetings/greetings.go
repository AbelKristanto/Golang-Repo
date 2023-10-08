package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Unit Testing

// Returned Name
func Hello(name string) (string, error) {
	// error faced
	if name == "" {
		return "", errors.New("Empty Name, Please Fill it!")
	}
	// return a greeting
	//message := fmt.Sprintf("Hi!, %v. Welcome home!", name)
	//return message, nil
	// create random name
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map associate names with message
	messages := make(map[string]string)

	// looping in the name, calling every using hello
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	// A slice of message format
	formats := []string{
		"Hi!, %v. Welcome Home!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// randomize
	return formats[rand.Intn(len(formats))]
}

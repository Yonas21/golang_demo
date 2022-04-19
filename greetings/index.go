package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// If no names is given
	if name == "" {
		return "", errors.New("Empty Name")
	}
	//return the message that embeds your name in a message

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos return a map that associates each of the name people
// with a greeting message

func Hellos(names []string) (map[string]string, error) {
	// map to associate names with messages
	messages := make(map[string]string)

	// loop through received slices of names
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

//set initial value for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message format
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"What's Hand in, %v",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

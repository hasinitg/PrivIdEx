package main

import "fmt"

type JSONDecodingError struct {
	ErroredMessage string
}

func (err JSONDecodingError) Error() string {
	return fmt.Sprintf("Could not decode the message: %v", err.ErroredMessage)
}
